package cmd

import (
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "sisco",
		Short: "Lightweight Service Discovery",
		Long: `sisco is a small and lightweight server providing the possibility to register services and
to query for them.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.sisco.yaml)")
}
