package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"sisco/internal/cfg"
	"sisco/internal/grpc/client"
)

var (
	bearerToken  string
	isAdminToken bool
)

func init() {
	adminCmd.AddCommand(adminLoginCmd)

	rootCmd.AddCommand(adminCmd)
}

var adminLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to sisco",
	Long:  `Login to Sisco gRPC-based administration interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminLogin(cmd, args)
	},
}

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Administrate sisco",
	Long:  `Sisco gRPC-based administration interface.`,
}

func execAdminLogin(cmd *cobra.Command, args []string) {
	var err error

	if len(args) != 2 {
		log.Fatalln(cmd.Usage())
	}

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	bearerToken, isAdminToken, err = client.GRPCLogin(listenAddr, args[0], args[1])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Token: %s, IsAdminToken: %t\n", bearerToken, isAdminToken)
}
