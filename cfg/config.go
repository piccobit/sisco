package cfg

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configuration contains the configuration of the app.
// Please adjust it to your own needs!
type configuration struct {
	Debug               bool     `yaml:"debug,omitempty"`
	GinReleaseMode      bool     `yaml:"ginReleaseMode,omitempty"`
	Port                int      `yaml:"port"`
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
	Config     configuration
	configName = ".sisco" // configName contains the name (without extension) of the configuration file.
	configType = "yaml"   // configType defines the type of the configuration file.
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
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
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
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into configuration struct, %v", err)
	}
}
