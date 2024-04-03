package ecr_repositories

import (
	"fmt"
	"go.dfds.cloud/ticli/cmds/configuration"
	"os"

	"go.dfds.cloud/ticli/selfservice"

	"github.com/spf13/cobra"
)

var (
	selfserviceClient *selfservice.SelfServiceClient
	description       string
)

var ECRCmd = &cobra.Command{
	Use:   "ecr",
	Short: "ecr-repo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Show ecr repositories")
	},
}

func InitializeECR(accessToken string) {
	queryCmd.PersistentFlags().StringP("id", "i", "", "capability id (currently not effective)")
	ECRCmd.AddCommand(queryCmd, createECRCmd)

	createECRCmd.PersistentFlags().StringVar(&description, "description", "", "adds a description to a ecr (required)")
	configuration.BindFlag("description", createECRCmd.PersistentFlags().Lookup("description"))

	selfserviceClient = selfservice.NewSelfServiceClient(accessToken)
}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query the API",
	Run: func(cmd *cobra.Command, args []string) {
		idInput, _ := cmd.Flags().GetString("id")
		fmt.Println(idInput)
		repositories := selfserviceClient.GetECRRepositories()
		fmt.Println(repositories)
	},
}

var createECRCmd = &cobra.Command{
	Use:   "create [NAME]",
	Short: "creates a new ecr repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Missing id")
			os.Exit(1)
		}
		ecrStruct := selfservice.CapabilityRequest{
			Name:        args[0],
			Description: description,
		}

		ecr := selfserviceClient.CreateECRRepo(ecrStruct)
		fmt.Println(ecr)

	},
}
