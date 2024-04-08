package kafka_topics

import (
	"go.dfds.cloud/ticli/cmds/outputwriter"
	"go.dfds.cloud/ticli/selfservice"

	"github.com/spf13/cobra"
)

var (
	selfserviceClient *selfservice.SelfServiceClient
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

	selfserviceClient = selfservice.NewSelfServiceClient(accessToken)

}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query the API",
	Run: func(cmd *cobra.Command, args []string) {
		topics := selfserviceClient.GetTopics()
		outputwriter.GetWriter().WriteData(topics)
	},
}
