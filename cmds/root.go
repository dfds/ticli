package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.dfds.cloud/ticli/cmds/authentication"
	"go.dfds.cloud/ticli/cmds/capability"
	"go.dfds.cloud/ticli/cmds/configuration"
	ecr_repositories "go.dfds.cloud/ticli/cmds/ecr-repositories"
	kafka_topics "go.dfds.cloud/ticli/cmds/kafka-topics"
	"go.dfds.cloud/ticli/cmds/outputwriter"
)

var (
	// Used for flags.
	cfgFile        string
	verbose        bool
	accessToken    string
	noVersionCheck bool
	outputFormat   string = "json"
)

var rootCmd = &cobra.Command{
	Use:     "ticli",
	Version: Version,
	Short:   "cli for dfds selfservice",
	Long:    `cli for dfds selfservice`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !noVersionCheck {
			remoteVersionCheck()
		}
		outputwriter.SetWriter(outputFormat)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Read configuration from files
	configuration.InitializeConfiguration()

	// Persistent Flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	configuration.BindFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "print moar stuff (default: false)")
	configuration.BindFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.PersistentFlags().StringVarP(&accessToken, "access-token", "t", "", "Provide an access token (or simply run authentication)")
	configuration.BindFlag("access-token", rootCmd.PersistentFlags().Lookup("access-token"))

	rootCmd.PersistentFlags().BoolVarP(&noVersionCheck, "no-version-check", "", false, "Disable version check")
	configuration.BindFlag("no-version-check", rootCmd.PersistentFlags().Lookup("no-version-check"))

	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", "", "Choose output format [json] (default: json)")
	configuration.BindFlag("output-format", rootCmd.PersistentFlags().Lookup("output-format"))

	// Commands
	capability.InitializeCapability(configuration.GetString("access-token"))
	rootCmd.AddCommand(capability.CapabilityCmd)

	ecr_repositories.InitializeECR(configuration.GetString("access-token"))
	rootCmd.AddCommand(ecr_repositories.ECRCmd)

	authentication.InitializeAuthentication()
	rootCmd.AddCommand(authentication.AuthenticationCmd)

	kafka_topics.InitTopics(configuration.GetString("access-token"))
	rootCmd.AddCommand(kafka_topics.TopicsCmd)
}
