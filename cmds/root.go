package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.dfds.cloud/mcdonalds/cmds/authentication"
	"go.dfds.cloud/mcdonalds/cmds/capability"
	"go.dfds.cloud/mcdonalds/cmds/configuration"
	ecr_repositories "go.dfds.cloud/mcdonalds/cmds/ecr-repositories"
	kafka_topics "go.dfds.cloud/mcdonalds/cmds/kafka-topics"
)

var (
	// Used for flags.
	cfgFile     string
	verbose     bool
	accessToken string
)

func main() {
	Execute()

}

var rootCmd = &cobra.Command{
	Use:   "mcdonalds",
	Short: "cli for querying dfds capabilities",
	Long:  `cli for querying dfds capabilities`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("I like burgers")
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

	rootCmd.PersistentFlags().StringVarP(&accessToken, "access-token", "", "", "Provide an access token (or simply run authentication)")
	configuration.BindFlag("access-token", rootCmd.PersistentFlags().Lookup("access-token"))

	// Commands
	rootCmd.AddCommand(configuration.TestConfigurationCmd)

	capability.InitializeCapability(configuration.GetString("access-token"))
	rootCmd.AddCommand(capability.CapabilityCmd)

	ecr_repositories.InitializeECR(configuration.GetString("access-token"))
	rootCmd.AddCommand(ecr_repositories.ECRCmd)

	authentication.InitializeAuthentication()
	rootCmd.AddCommand(authentication.AuthenticationCmd)

	kafka_topics.InitTopics(configuration.GetString("access-token"))
	rootCmd.AddCommand(kafka_topics.TopicsCmd)

	rootCmd.AddCommand(configuration.TestConfigurationCmd)
}
