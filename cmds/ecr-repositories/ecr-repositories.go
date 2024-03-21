package ecr_repositories

import (
	"fmt"

	"go.dfds.cloud/mcdonalds/selfservice"

	"github.com/spf13/cobra"
)

var (
	selfserviceClient *selfservice.SelfServiceClient
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
	ECRCmd.AddCommand(queryCmd)

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
