package cmd

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx"
	"github.com/spf13/cobra"
	"sisco/internal/exit"
	"sisco/internal/rpc/crpc"
	"sisco/internal/utils"
)

var listCmd = &cobra.Command{
	Use:   "list [command]",
	Short: "List components",
	Long:  `List the specified components.`,
}

var listServiceCmd = &cobra.Command{
	Use:   "service [service name] [area name]",
	Short: "List service in area",
	Long:  `List specified service in the specified area.`,
	Run: func(cmd *cobra.Command, args []string) {
		execListService(cmd, args)
	},
}

var listServicesCmd = &cobra.Command{
	Use:   "services [area name]",
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

var listTagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "List tags",
	Long:  `List all tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		execListTags(cmd, args)
	},
}

func init() {
	listCmd.AddCommand(listServiceCmd)
	listCmd.AddCommand(listServicesCmd)
	listCmd.AddCommand(listAreasCmd)
	listCmd.AddCommand(listTagsCmd)

	listCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Auth token")

	rootCmd.AddCommand(listCmd)
}

func execListService(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		exit.Fatalln(1, cmd.Usage())
	}

	rpcClient, err := crpc.Default()
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	l, err := rpcClient.ListServiceInArea(getToken(), args[0], args[1])
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	fmt.Println(utils.JSONify(StatusCode{"OK", l}, pretty))
}

func execListServices(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		exit.Fatalln(1, cmd.Usage())
	}

	rpcClient, err := crpc.Default()
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

	rpcClient, err := crpc.Default()
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	l, err := rpcClient.ListAreas(getToken())
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	fmt.Println(utils.JSONify(StatusCode{"OK", l}, pretty))
}

func execListTags(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		exit.Fatalln(1, cmd.Usage())
	}

	rpcClient, err := crpc.Default()
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	l, err := rpcClient.ListTags(getToken())
	if err != nil {
		exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", err.Error()}, pretty))
	}

	fmt.Println(utils.JSONify(StatusCode{"OK", l}, pretty))
}
