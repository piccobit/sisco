package cfg

import (
	"log"
	"os"

	"github.com/a8m/envsubst"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Configuration contains the Configuration of the app.
// Please adjust it to your own needs!
type Configuration struct {
	Debug               bool     `yaml:"debug,omitempty"`
	GinReleaseMode      bool     `yaml:"ginReleaseMode,omitempty"`
	Port                int      `yaml:"port"`
	GRPCPort            int      `yaml:"gRPCPort"`
	UseTLS              bool     `yaml:"useTLS"`
	TLSCertFile         string   `yaml:"tlsCertFile"`
	TLSKeyFile          string   `yaml:"tlsKeyFile"`
	DBType              string   `yaml:"dbType"`
	DBHost              string   `yaml:"dbHost"`
	DBPort              int      `yaml:"dbPort"`
	DBName              string   `yaml:"dbName"`
	DBUser              string   `yaml:"dbUser"`
	DBPassword          string   `yaml:"dbPassword"`
	DBSSLMode           string   `yaml:"dbSSLMode"`
	TrustedProxies      []string `yaml:"trustedProxies,omitempty"`
	LdapURL             string   `yaml:"ldapURL"`
	LdapBaseDN          string   `yaml:"ldapBaseDN"`
	LdapBindDN          string   `yaml:"ldapBindDN"`
	LdapBindPassword    string   `yaml:"ldapBindPassword"`
	LdapFilterUsersDN   string   `yaml:"ldapFilterUsersDN"`
	LdapFilterAdminsDN  string   `yaml:"ldapFilterAdminsDN"`
	TokenValidInSeconds int      `yaml:"tokenValidInSeconds"`
}

var (
	Config     Configuration
	configName = ".sisco" // configName contains the name (without extension) of the Configuration file.
	configType = "yaml"   // configType defines the type of the Configuration file.
	configFile string
)

func init() {
	// cobra.OnInitialize(initConfig)
}

func New(cfgFile string) {
	configFile = cfgFile

	initConfig()
}

func initConfig() {
	if configFile != "" {
		filename, err := homedir.Expand(configFile)
		if err != nil {
			log.Fatalf("home directory for config file path '%s' not found: %v", configFile, err)
		}

		filename, err = envsubst.String(filename)
		if err != nil {
			log.Fatalf("could not envsubst config file path '%s' not found: %v", configFile, err)
		}

		// Use config file from the flag.
		viper.SetConfigFile(filename)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".sisco" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("unable to read configuration file, %v", err)
	}

	if viper.GetBool("debug") {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into configuration struct, %v", err)
	}
}
