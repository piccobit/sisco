package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sisco/internal/cfg"
	"sisco/internal/crpc"
	"sisco/internal/exit"
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
	Use:   "register-area <area> <description>",
	Short: "Register area",
	Long:  `Register a new area.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminRegisterArea(cmd, args)
	},
}

var adminRegisterServiceCmd = &cobra.Command{
	Use:   "register-service <service> <area> <description> <protocol> <host> <port> <tag-1> ... <tag-n>",
	Short: "Register service",
	Long:  `Register a new service in an area.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminRegisterService(cmd, args)
	},
}

var adminDeleteAreaCmd = &cobra.Command{
	Use:   "delete-area <area>",
	Short: "Delete area",
	Long:  `Delete an area. The area must be empty or an error is issued.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminDeleteArea(cmd, args)
	},
}

var adminDeleteServiceCmd = &cobra.Command{
	Use:   "delete-service <service> <area>",
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
		exit.Fatalln(1, cmd.Usage())
	}

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	rpcClient, err := crpc.New(
		crpc.ListenAddr(listenAddr),
		crpc.UseTLS(cfg.Config.UseTLS),
		crpc.TLSCertFile(cfg.Config.TLSCertFile),
	)
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	err = rpcClient.RegisterArea(getToken(), args[0], args[1])
	if err == nil {
		fmt.Println(utils.JSONify(StatusCode{"OK", ""}, pretty))
	} else {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}
}

func execAdminRegisterService(cmd *cobra.Command, args []string) {
	if len(args) < 6 {
		exit.Fatalln(1, cmd.Usage())
	}

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	grpcClient, err := crpc.New(
		crpc.ListenAddr(listenAddr),
		crpc.UseTLS(cfg.Config.UseTLS),
		crpc.TLSCertFile(cfg.Config.TLSCertFile),
	)
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	err = grpcClient.RegisterService(getToken(), args[0], args[1], args[2], args[3], args[4], args[5], args[6:]...)
	if err == nil {
		fmt.Println(utils.JSONify(StatusCode{"OK", ""}, pretty))
	} else {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}
}

func execAdminDeleteArea(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		exit.Fatalln(1, cmd.Usage())
	}

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	grpcClient, err := crpc.New(
		crpc.ListenAddr(listenAddr),
		crpc.UseTLS(cfg.Config.UseTLS),
		crpc.TLSCertFile(cfg.Config.TLSCertFile),
	)
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	err = grpcClient.DeleteArea(getToken(), args[0])
	if err == nil {
		fmt.Println(utils.JSONify(StatusCode{"OK", ""}, pretty))
	} else {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}
}

func execAdminDeleteService(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		exit.Fatalln(1, cmd.Usage())
	}

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	grpcClient, err := crpc.New(
		crpc.ListenAddr(listenAddr),
		crpc.UseTLS(cfg.Config.UseTLS),
		crpc.TLSCertFile(cfg.Config.TLSCertFile),
	)
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	err = grpcClient.DeleteService(getToken(), args[0], args[1])
	if err == nil {
		fmt.Println(utils.JSONify(StatusCode{Status: "OK"}, pretty))
	} else {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}
}
