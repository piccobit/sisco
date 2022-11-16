package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"sisco/internal/cfg"
	"sisco/internal/crpc"
	"sisco/internal/utils"
)

type StatusCode struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
}
type AuthTokenInfo struct {
	Token        string `json:"token"`
	IsAdminToken bool   `json:"isAdminToken"`
}

func init() {
	adminCmd.AddCommand(adminRegisterAreaCmd)
	adminCmd.AddCommand(adminRegisterServiceCmd)
	adminCmd.AddCommand(adminDeleteAreaCmd)
	adminCmd.AddCommand(adminDeleteServiceCmd)

	rootCmd.AddCommand(adminCmd)
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

func execAdminRegisterArea(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatalln(cmd.Usage())
	}

	getToken()

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	rpcClient, err := crpc.New(
		crpc.ListenAddr(listenAddr),
		crpc.UseTLS(cfg.Config.UseTLS),
		crpc.TLSCertFile(cfg.Config.TLSCertFile),
	)
	if err != nil {
		log.Fatalln(utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	err = rpcClient.RegisterArea(token, args[0], args[1])
	if err == nil {
		log.Println(utils.JSONify(StatusCode{"OK", ""}, pretty))
	} else {
		log.Fatalln(utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}
}

func execAdminRegisterService(cmd *cobra.Command, args []string) {
	if len(args) < 6 {
		log.Fatalln(cmd.Usage())
	}

	getToken()

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	grpcClient, err := crpc.New(
		crpc.ListenAddr(listenAddr),
		crpc.UseTLS(cfg.Config.UseTLS),
		crpc.TLSCertFile(cfg.Config.TLSCertFile),
	)
	if err != nil {
		log.Fatalln(utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	err = grpcClient.RegisterService(token, args[0], args[1], args[2], args[3], args[4], args[5], args[6:]...)
	if err == nil {
		log.Println(utils.JSONify(StatusCode{"OK", ""}, pretty))
	} else {
		log.Fatalln(utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}
}

func execAdminDeleteArea(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatalln(cmd.Usage())
	}

	getToken()

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	grpcClient, err := crpc.New(
		crpc.ListenAddr(listenAddr),
		crpc.UseTLS(cfg.Config.UseTLS),
		crpc.TLSCertFile(cfg.Config.TLSCertFile),
	)
	if err != nil {
		log.Fatalln(utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	err = grpcClient.DeleteArea(args[0], args[1])
	if err == nil {
		log.Println(utils.JSONify(StatusCode{"OK", ""}, pretty))
	} else {
		log.Fatalln(utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}
}

func execAdminDeleteService(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		log.Fatalln(cmd.Usage())
	}

	getToken()

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	grpcClient, err := crpc.New(
		crpc.ListenAddr(listenAddr),
		crpc.UseTLS(cfg.Config.UseTLS),
		crpc.TLSCertFile(cfg.Config.TLSCertFile),
	)
	if err != nil {
		log.Fatalln(utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	err = grpcClient.DeleteService(args[0], args[1], args[2])
	if err == nil {
		log.Println(utils.JSONify(StatusCode{Status: "OK"}, pretty))
	} else {
		log.Fatalln(utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}
}
