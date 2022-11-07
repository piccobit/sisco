package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"sisco/internal/cfg"
	"sisco/internal/grpc/client"
)

func init() {
	adminCmd.AddCommand(adminLoginCmd)
	adminCmd.AddCommand(adminRegisterAreaCmd)
	adminCmd.AddCommand(adminDeleteAreaCmd)

	rootCmd.AddCommand(adminCmd)
}

var adminLoginCmd = &cobra.Command{
	Use:   "login <user> <password>",
	Short: "Login to sisco",
	Long:  `Login to Sisco gRPC-based administration interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminLogin(cmd, args)
	},
}

var adminRegisterAreaCmd = &cobra.Command{
	Use:   "register-area <auth-token> <name> <description>",
	Short: "Register area",
	Long:  `Register a new area.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminRegisterArea(cmd, args)
	},
}

var adminDeleteAreaCmd = &cobra.Command{
	Use:   "delete-area <auth-token> <name> <description>",
	Short: "Delete area",
	Long:  `Delete an area. The area must be empty or an error is issued.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminDeleteArea(cmd, args)
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

	bearerToken, isAdminToken, err := client.Login(listenAddr, args[0], args[1])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Token: %s, IsAdminToken: %t\n", bearerToken, isAdminToken)
}

func execAdminRegisterArea(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		log.Fatalln(cmd.Usage())
	}

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	err := client.RegisterArea(listenAddr, args[0], args[1], args[2])
	if err == nil {
		fmt.Println("Status: OK")
	} else {
		fmt.Printf("Status: NOT OK - %v\n", err)
	}
}

func execAdminDeleteArea(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatalln(cmd.Usage())
	}

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	err := client.DeleteArea(listenAddr, args[0], args[1])
	if err == nil {
		fmt.Println("Status: OK")
	} else {
		fmt.Printf("Status: NOT OK - %v\n", err)
	}
}
