package cmd

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx"
	"github.com/spf13/cobra"
	"log"
	"sisco/internal/cfg"
	"sisco/internal/crpc"
	"sisco/internal/utils"
)

var listCmd = &cobra.Command{
	Use:   "list <command>",
	Short: "List components",
	Long:  `List the specified components.`,
}

var listServicesCmd = &cobra.Command{
	Use:   "services <token> <area-name>",
	Short: "List services",
	Long:  `List all services in a specified area.`,
	Run: func(cmd *cobra.Command, args []string) {
		execListServices(cmd, args)
	},
}

var listAreasCmd = &cobra.Command{
	Use:   "areas <token>",
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

	l, err := rpcClient.ListServices(token, args[0])
	if err != nil {
		log.Fatalln(utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	log.Println(utils.JSONify(StatusCode{"OK", l}, pretty))
}

func execListAreas(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
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

	l, err := rpcClient.ListAreas(token)
	if err != nil {
		log.Fatalln(utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	log.Println(utils.JSONify(StatusCode{"OK", l}, pretty))
}
