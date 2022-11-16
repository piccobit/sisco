package cmd

import (
	atlas "ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqltool"
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"path/filepath"
	entMigrate "sisco/ent/migrate"
	"sisco/internal/cfg"
	"sisco/internal/exit"
	"strconv"
)

var (
	migrationsBaseDir string
)

var migrateCmd = &cobra.Command{
	Use:   "migrate <command>",
	Short: "Support database migration",
	Long:  `Supports the database migration using files in the migration directory.`,
}

var migrateGenerateCmd = &cobra.Command{
	Use:   "generate <dev-database-name> <migration-name>",
	Short: "Generate database migration files",
	Long: `Generates the database migration files in the base migration directory using the provided
empty 'dev'' database. The 'dev' database shall use the same permissions as the productive database.`,
	Run: func(cmd *cobra.Command, args []string) {
		execMigrateGenerate(cmd, args)
	},
}

var migrateApplyCmd = &cobra.Command{
	Use:   "apply [<num-migrations>]",
	Short: "Apply database migration files",
	Long:  `Apply the specified number of migration files in the database migrations directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		execMigrateApply(args)
	},
}

func init() {
	migrateCmd.AddCommand(migrateGenerateCmd)
	migrateCmd.AddCommand(migrateApplyCmd)

	migrateCmd.PersistentFlags().StringVarP(&migrationsBaseDir, "migrations-base-dir", "m", "migrations", "database migrations base directory")
	rootCmd.AddCommand(migrateCmd)
}

func execMigrateGenerate(cmd *cobra.Command, args []string) {
	var err error

	dbType := cfg.Config.DBType

	if dbType == "mariadb" {
		dbType = "mysql"
	}

	ctx := context.Background()

	var dir *atlas.LocalDir

	// Create a local migration directory able to understand Atlas migration file format for replay.
	if migrationsBaseDir == "migrations" {
		cwd, err := os.Getwd()
		if err != nil {
			exit.Fatalf(1, "failed getting current working directory: %v", err)
		}

		migrationsDir := filepath.Join(cwd, "migrations", dbType)

		checkMigrationsDir(migrationsDir)

		dir, err = atlas.NewLocalDir(migrationsDir)
	} else {
		migrationsDir, err := filepath.Abs(filepath.Join(migrationsBaseDir, dbType))
		if err != nil {
			exit.Fatalf(1, "failed getting absolute path: %v", err)
		}

		checkMigrationsDir(migrationsDir)

		dir, err = atlas.NewLocalDir(migrationsDir)
	}

	if err != nil {
		exit.Fatalf(1, "failed creating atlas migration directory: %v", err)
	}

	var dbDialect string

	switch dbType {
	case "postgres":
		dbDialect = dialect.Postgres

	case "mysql":
		dbDialect = dialect.MySQL
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dbDialect),               // Ent dialect to use
		schema.WithFormatter(sqltool.GolangMigrateFormatter),
		schema.WithForeignKeys(false),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	}

	if len(args) != 2 {
		exit.Fatalln(1, cmd.Usage())
	}

	// We're computing the diff by connection to the provided empty dev database.
	// This dev database should use the same permissions as the production database.
	var dbURL string

	switch dbType {
	case "postgres":
		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.Config.DBUser,
			cfg.Config.DBPassword,
			cfg.Config.DBHost,
			cfg.Config.DBPort,
			args[0],
			cfg.Config.DBSSLMode,
		)

	case "mysql":
		dbURL = fmt.Sprintf("mysql://%s:%s@%s:%d/%s",
			cfg.Config.DBUser,
			cfg.Config.DBPassword,
			cfg.Config.DBHost,
			cfg.Config.DBPort,
			args[0],
		)
	}

	// Generate migrations using Atlas support for Postgres (note the Ent dialect option passed above).
	err = entMigrate.NamedDiff(ctx, dbURL, args[1], opts...)
	if err != nil {
		exit.Fatalf(1, "failed generating migration files: %v", err)
	}
}

func execMigrateApply(args []string) {
	var err error

	dbType := cfg.Config.DBType

	if dbType == "mariadb" {
		dbType = "mysql"
	}

	numMigrations := 0

	if len(args) != 0 {
		numMigrations, err = strconv.Atoi(args[0])
		if err != nil {
			exit.Fatalf(1, "could not get number of migrations: %v", err)
		}
	}
	var dbURL string

	switch dbType {
	case "postgres":
		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.Config.DBUser,
			cfg.Config.DBPassword,
			cfg.Config.DBHost,
			cfg.Config.DBPort,
			cfg.Config.DBName,
			cfg.Config.DBSSLMode,
		)

	case "mysql":
		dbURL = fmt.Sprintf("mysql://%s:%s@%s:%d/%s",
			cfg.Config.DBUser,
			cfg.Config.DBPassword,
			cfg.Config.DBHost,
			cfg.Config.DBPort,
			cfg.Config.DBName,
		)
	}

	db, err := sql.Open(dbType, dbURL)
	if err != nil {
		exit.Fatalf(1, "failed opening database connection: %v", err)
	}

	var driver database.Driver

	switch dbType {
	case "postgres":
		driver, err = postgres.WithInstance(db, &postgres.Config{})
	case "mariadb":
		fallthrough
	case "mysql":
		driver, err = mysql.WithInstance(db, &mysql.Config{})
	}

	if err != nil {
		exit.Fatalf(1, "failed getting database driver: %v", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		exit.Fatalf(1, "failed getting current working directory: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", filepath.Join(cwd, "migrations", dbType)),
		dbType, driver)
	if err != nil {
		exit.Fatalf(1, "failed setting up migration instance: %v", err)
	}

	if numMigrations != 0 {
		err = m.Steps(numMigrations)
	} else {
		err = m.Up()
	}
	if err != nil {
		exit.Fatalf(1, "database migration failed: %v", err)
	}
}

func checkMigrationsDir(migrationsDir string) {
	fileInfo, err := os.Stat(migrationsDir)
	if errors.Is(err, fs.ErrNotExist) {
		exit.Fatalf(1, "migrations directory '%s' does not exist: %v", migrationsDir, err)
	}

	if !fileInfo.IsDir() {
		exit.Fatalf(1, "migrations directory '%s' is a file", migrationsDir)
	}

	file, err := os.Open(migrationsDir)
	if err != nil {
		exit.Fatalf(1, "could not open migrations directory '%s': %v", migrationsDir, err)
	}

	dirEntries, err := file.Readdirnames(0)
	if err != nil {
		exit.Fatalf(1, "could not read migrations directory '%s': %v", migrationsDir, err)
	}

	if len(dirEntries) > 0 {
		exit.Fatalf(1, "migrations directory '%s' is not empty", migrationsDir)
	}
}
