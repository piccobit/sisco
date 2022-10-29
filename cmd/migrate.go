package cmd

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/spf13/cobra"

	"sisco/ent/migrate"

	_ "github.com/jackc/pgx"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate <command>",
	Short: "Support database migration",
	Long:  `Supports the database migration using files in the migration directory.`,
}

var migrateGenerateCmd = &cobra.Command{
	Use:   "generate <name>",
	Short: "Generate database migration files",
	Long:  `Generates the database migration files in the migration directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		execMigrateGenerate(args)
	},
}

var migrateApplyCmd = &cobra.Command{
	Use:   "apply [<num-migrations>]",
	Short: "Apply database migration files",
	Long:  `Apply the specified number of database migration files in the migration directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		execMigrateApply(args)
	},
}

func init() {
	migrateCmd.AddCommand(migrateGenerateCmd)
	migrateCmd.AddCommand(migrateApplyCmd)

	rootCmd.AddCommand(migrateCmd)
}

func execMigrateGenerate(args []string) {
	ctx := context.Background()

	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	if len(args) != 1 {
		log.Fatalln("migration name is required. Use: 'sisco migrate generate <name>'")
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		Config.DBUser,
		Config.DBPassword,
		Config.DBHost,
		Config.DBPort,
		Config.DBName,
		Config.DBSSLMode,
	)

	// Generate migrations using Atlas support for Postgres (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, dbURL, args[0], opts...)
	if err != nil {
		log.Fatalf("failed generating migration files: %v", err)
	}
}

func execMigrateApply(args []string) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		Config.DBUser,
		Config.DBPassword,
		Config.DBHost,
		Config.DBPort,
		Config.DBName,
		Config.DBSSLMode,
	)

	var execArgs []string

	if len(args) != 1 {
		execArgs = []string{
			"migrate",
			"apply",
			"--url",
			dbURL,
		}
	} else {
		execArgs = []string{
			"migrate",
			"apply",
			"--url",
			dbURL,
			"--latest",
			args[0],
		}
	}

	out, err := exec.Command("atlas", execArgs...).Output()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(string(out))
}
