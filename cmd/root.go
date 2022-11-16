package cmd

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"sisco/internal/cfg"
	"sisco/internal/exit"
	"sisco/internal/rpc/crpc"
	"sisco/internal/utils"
)

var (
	debug   bool
	pretty  bool
	cfgFile string
	token   string
)

var rootCmd = &cobra.Command{
	Use:   "sisco",
	Short: "Lightweight ServiceExtended Discovery",
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
		exit.Fatalln(1, "could not bind to flag '--debug'")
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
		exit.Fatalln(1, cmd.Usage())
	}

	grpcClient, err := crpc.Default()
	if err != nil {
		exit.Fatalln(1, err)
	}

	bearerToken, isAdminToken, err := grpcClient.Login(args[0], args[1])
	if err != nil {
		exit.Fatalln(1, err)
	}

	fmt.Println(utils.JSONify(AuthTokenInfo{
		Token:        bearerToken,
		IsAdminToken: isAdminToken,
	}, pretty))
}

func getToken() string {
	if len(token) == 0 {
		var ok bool

		token, ok = os.LookupEnv("SISCO_TOKEN")
		if !ok {
			exit.Fatalln(1, utils.JSONify(StatusCode{"NOT OK", "missing token"}, pretty))
		}
	}

	return token
}
