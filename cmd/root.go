package cmd

import (
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sisco/cfg"
)

var (
	debug   bool
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "sisco",
	Short: "Lightweight Service Discovery",
	Long: `sisco is a small and lightweight server providing the possibility to register services and
to query for them.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cfg.New(cfgFile)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug output")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.sisco.yaml)")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
