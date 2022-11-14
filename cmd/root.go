package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sisco/internal/cfg"
	"sisco/internal/grpc/client"
)

var (
	debug   bool
	pretty  bool
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "sisco",
	Short: "Lightweight Service Discovery",
	Long: `sisco is a small and lightweight server providing the possibility to register services and
to query for them.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cfg.New(cfgFile)
	},
}

var loginCmd = &cobra.Command{
	Use:   "login <user> <password>",
	Short: "Login to sisco",
	Long:  `Login to Sisco gRPC-based administration interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		execLogin(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug output")
	rootCmd.PersistentFlags().BoolVarP(&pretty, "pretty", "p", false, "enable pretty output")
	err := viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	if err != nil {
		log.Fatalln("could not bind to flag '--debug'")
	}

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.sisco.yaml)")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func execLogin(cmd *cobra.Command, args []string) {
	var err error

	if len(args) != 2 {
		log.Fatalln(cmd.Usage())
	}

	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	grpcClient, err := client.New(listenAddr)
	if err == nil {
		fmt.Println(StatusCode{"OK", ""})
	} else {
		fmt.Println(StatusCode{"NOT OK", err.Error()})
	}

	bearerToken, isAdminToken, err := grpcClient.Login(args[0], args[1])
	if err != nil {
		log.Fatalln(err)
	}

	jsonBlob, err := json.Marshal(AuthTokenInfo{
		Token:        bearerToken,
		IsAdminToken: isAdminToken,
	})

	fmt.Println(string(jsonBlob))
}
