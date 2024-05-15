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
	jsonOutput     bool
	yamlOutput     bool
)

func setOutputWriter() {
	if jsonOutput && yamlOutput {
		outputwriter.GetWriter().WriteError(fmt.Errorf("only one of json or yaml output can be selected"))
		os.Exit(1)
	}

	if jsonOutput {
		outputwriter.SetWriter(outputwriter.CreateJsonWriter())
	}
	if yamlOutput {
		outputwriter.SetWriter(outputwriter.CreateYamlWriter())
	}
}

var rootCmd = &cobra.Command{
	Use:     "ticli",
	Version: Version,
	Short:   "cli for dfds selfservice",
	Long:    `cli for dfds selfservice`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		setOutputWriter()
		if !noVersionCheck {
			remoteVersionCheck()
		}
	},
}

func Execute() {
	outputwriter.SetWriter(outputwriter.CreateDefaultWriter())
	if err := rootCmd.Execute(); err != nil {
		outputwriter.GetWriter().WriteError(fmt.Errorf("error while executing root command %s", err))
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

	rootCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "", false, "Json output")
	configuration.BindFlag("json", rootCmd.PersistentFlags().Lookup("json"))

	rootCmd.PersistentFlags().BoolVarP(&yamlOutput, "yaml", "", false, "Yaml output")
	configuration.BindFlag("yaml", rootCmd.PersistentFlags().Lookup("yaml"))

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
