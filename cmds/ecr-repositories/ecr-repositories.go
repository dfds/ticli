package ecr_repositories

import (
	"fmt"
	"os"

	"go.dfds.cloud/ticli/cmds/configuration"

	"go.dfds.cloud/ticli/cmds/outputwriter"
	"go.dfds.cloud/ticli/selfservice"

	"github.com/spf13/cobra"
)

var (
	selfserviceClient *selfservice.SelfServiceClient
	description       string
)

var ECRCmd = &cobra.Command{
	Use:   "ecr",
	Short: "Manage Selfservice AWS ECR repositories",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func InitializeECR(accessToken string) {
	queryCmd.PersistentFlags().StringP("id", "i", "", "capability id (currently not effective)")
	ECRCmd.AddCommand(queryCmd, createECRCmd)

	createECRCmd.PersistentFlags().StringVar(&description, "description", "", "adds a description to a ecr (required)")
	configuration.BindFlag("description", createECRCmd.PersistentFlags().Lookup("description"))
	ECRCmd.AddCommand(queryCmd)

	selfserviceClient = selfservice.NewSelfServiceClient(accessToken)
}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query the API",
	Run: func(cmd *cobra.Command, args []string) {
		repositories := selfserviceClient.GetECRRepositories()
		outputwriter.GetWriter().WriteData(repositories)
	},
}

var createECRCmd = &cobra.Command{
	Use:   "create [id]",
	Short: "creates a new ecr repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			outputwriter.GetWriter().WriteError(fmt.Errorf("missing id"))
			os.Exit(1)
		}
		ecrStruct := selfservice.CapabilityRequest{
			Name:        args[0],
			Description: description,
		}

		ecr := selfserviceClient.CreateECRRepo(ecrStruct)
		outputwriter.GetWriter().WriteData(ecr)

	},
}
