package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sisco/internal/auth"
	"sisco/internal/exit"
	"sisco/internal/rpc/crpc"
	"sisco/internal/utils"
)

type StatusCode struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
}

type AuthTokenInfo struct {
	Token       string           `json:"token"`
	Permissions auth.Permissions `json:"permissions"`
}

var adminRegisterCmd = &cobra.Command{
	Use:   "register [components]",
	Short: "Register new components",
	Long:  `Register the specified components.`,
}

var adminRegisterAreaCmd = &cobra.Command{
	Use:   "area [area name] [description]",
	Short: "Register area",
	Long:  `Register a new area.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminRegisterArea(cmd, args)
	},
}

var adminRegisterServiceCmd = &cobra.Command{
	Use:   "service [service name] [area name] [description] [protocol] [host] [port] [tag-1] ... [tag-n]",
	Short: "Register service",
	Long:  `Register a new service in an area.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminRegisterService(cmd, args)
	},
}

var adminDeleteCmd = &cobra.Command{
	Use:   "delete [components]",
	Short: "Delete a components",
	Long:  `Delete the specified components.`,
}

var adminDeleteAreaCmd = &cobra.Command{
	Use:   "area [area name]",
	Short: "Delete area",
	Long:  `Delete an area. The area must be empty or an error is issued.`,
	Run: func(cmd *cobra.Command, args []string) {
		execAdminDeleteArea(cmd, args)
	},
}

var adminDeleteServiceCmd = &cobra.Command{
	Use:   "service [service name] [area name]",
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

func init() {
	adminRegisterCmd.AddCommand(adminRegisterAreaCmd)
	adminRegisterCmd.AddCommand(adminRegisterServiceCmd)
	adminDeleteCmd.AddCommand(adminDeleteAreaCmd)
	adminDeleteCmd.AddCommand(adminDeleteServiceCmd)
	adminCmd.AddCommand(adminRegisterCmd)
	adminCmd.AddCommand(adminDeleteCmd)

	rootCmd.AddCommand(adminCmd)
}

func execAdminRegisterArea(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		exit.Fatalln(1, cmd.Usage())
	}

	rpcClient, err := crpc.Default()
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

	rpcClient, err := crpc.Default()
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	err = rpcClient.RegisterService(getToken(), args[0], args[1], args[2], args[3], args[4], args[5], args[6:]...)
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

	rpcClient, err := crpc.Default()
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	err = rpcClient.DeleteArea(getToken(), args[0])
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

	rpcClient, err := crpc.Default()
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	err = rpcClient.DeleteService(getToken(), args[0], args[1])
	if err == nil {
		fmt.Println(utils.JSONify(StatusCode{Status: "OK"}, pretty))
	} else {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}
}
