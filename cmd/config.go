package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type config struct {
	GinReleaseMode bool     `yaml:"ginReleaseMode"`
	Port           string   `yaml:"port"`
	DBHost         string   `yaml:"dbHost"`
	DBPort         string   `yaml:"dbPort"`
	DBName         string   `yaml:"dbName"`
	DBUser         string   `yaml:"dbUser"`
	DBPassword     string   `yaml:"dbPassword"`
	DBSSLMode      string   `yaml:"dbSSLMode"`
	TrustedProxies []string `yaml:"trustedProxies"`
}

var (
	Config config
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".sisco")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into config struct, %v", err)
	}
}
