package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"sisco/internal/cfg"
	"sisco/internal/grpc/client"
)

type StatusCode struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type AuthTokenInfo struct {
	Token        string `json:"token"`
	IsAdminToken bool   `json:"isAdminToken"`
}

func init() {
	adminCmd.AddCommand(adminLoginCmd)
	adminCmd.AddCommand(adminRegisterAreaCmd)
	adminCmd.AddCommand(adminRegisterServiceCmd)
	adminCmd.AddCommand(adminDeleteAreaCmd)
	adminCmd.AddCommand(adminDeleteServiceCmd)

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
	Use:   "register-area <auth-token> <area> <description>",
	Short: "Register area",
	Long:  `Register a new area.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminRegisterArea(cmd, args)
	},
}

var adminRegisterServiceCmd = &cobra.Command{
	Use:   "register-service <auth-token> <service> <area> <description> <protocol> <host> <port> <tag-1> ... <tag-n>",
	Short: "Register service",
	Long:  `Register a new service in an area.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminRegisterService(cmd, args)
	},
}

var adminDeleteAreaCmd = &cobra.Command{
	Use:   "delete-area <auth-token> <area>",
	Short: "Delete area",
	Long:  `Delete an area. The area must be empty or an error is issued.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminDeleteArea(cmd, args)
	},
}

var adminDeleteServiceCmd = &cobra.Command{
	Use:   "delete-service <auth-token> <service> <area>",
	Short: "Delete service",
	Long:  `Delete a service in an area.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminDeleteService(cmd, args)
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

	jsonBlob, err := json.Marshal(AuthTokenInfo{
		Token:        bearerToken,
		IsAdminToken: isAdminToken,
	})

	fmt.Println(string(jsonBlob))
}

func execAdminRegisterArea(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		log.Fatalln(cmd.Usage())
	}

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	err := client.RegisterArea(listenAddr, args[0], args[1], args[2])
	if err == nil {
		fmt.Println(StatusCode{"OK", ""})
	} else {
		fmt.Println(StatusCode{"NOT OK", err.Error()})
	}
}

func execAdminRegisterService(cmd *cobra.Command, args []string) {
	if len(args) < 7 {
		log.Fatalln(cmd.Usage())
	}

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	err := client.RegisterService(listenAddr, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7:]...)
	if err == nil {
		fmt.Println(StatusCode{"OK", ""})
	} else {
		fmt.Println(StatusCode{"NOT OK", err.Error()})
	}
}

func execAdminDeleteArea(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatalln(cmd.Usage())
	}

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	err := client.DeleteArea(listenAddr, args[0], args[1])
	if err == nil {
		fmt.Println(StatusCode{"OK", ""})
	} else {
		fmt.Println(StatusCode{"NOT OK", err.Error()})
	}
}

func execAdminDeleteService(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		log.Fatalln(cmd.Usage())
	}

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	err := client.DeleteService(listenAddr, args[0], args[1], args[2])
	if err == nil {
		fmt.Println(StatusCode{Status: "OK"})
	} else {
		fmt.Printf("Status: NOT OK - %v\n", err)
		fmt.Println(StatusCode{"NOT OK", err.Error()})
	}
}
