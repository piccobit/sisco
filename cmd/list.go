package cmd

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx"
	"github.com/spf13/cobra"
	"sisco/internal/cfg"
	"sisco/internal/crpc"
	"sisco/internal/exit"
	"sisco/internal/utils"
)

var listCmd = &cobra.Command{
	Use:   "list <command>",
	Short: "List components",
	Long:  `List the specified components.`,
}

var listServicesCmd = &cobra.Command{
	Use:   "services <area-name>",
	Short: "List services",
	Long:  `List all services in a specified area.`,
	Run: func(cmd *cobra.Command, args []string) {
		execListServices(cmd, args)
	},
}

var listAreasCmd = &cobra.Command{
	Use:   "areas",
	Short: "List areas",
	Long:  `List all areas.`,
	Run: func(cmd *cobra.Command, args []string) {
		execListAreas(cmd, args)
	},
}

func init() {
	listCmd.AddCommand(listServicesCmd)
	listCmd.AddCommand(listAreasCmd)

	listCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Auth token")

	rootCmd.AddCommand(listCmd)
}

func execListServices(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
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

	l, err := rpcClient.ListServices(getToken(), args[0])
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	fmt.Println(utils.JSONify(StatusCode{"OK", l}, pretty))
}

func execListAreas(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
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

	l, err := rpcClient.ListAreas(getToken())
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	fmt.Println(utils.JSONify(StatusCode{"OK", l}, pretty))
}
