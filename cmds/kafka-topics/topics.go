package kafka_topics

import (
	"context"

	"go.dfds.cloud/ticli/cmds/outputwriter"
	"go.dfds.cloud/ticli/openapiclient"
	"go.dfds.cloud/ticli/selfservice"

	"github.com/spf13/cobra"
)

var (
	selfserviceClient *openapiclient.ClientWithResponses
)

var TopicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "query topics",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func InitTopics(accessToken string) {
	TopicsCmd.AddCommand(queryCmd)

	selfserviceClient = selfservice.NewGeneratedClient(accessToken)

}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query the API",
	Run: func(cmd *cobra.Command, args []string) {

		topics, err := selfserviceClient.GetKafkatopicsWithResponse(context.Background(), nil)

		if err != nil {
			outputwriter.GetWriter().WriteError(err)
		}

		outputwriter.GetWriter().WriteData(topics.JSON200)
	},
}
