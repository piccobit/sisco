package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version string = "dev"
	Commit  string = "-"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of sisco",
	Long:  `All software has versions. This is sisco's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sisco - Lightweight Service Discovery -- %s -- %s\n", Version, Commit)
	},
}
