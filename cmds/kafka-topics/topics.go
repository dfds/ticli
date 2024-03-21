package kafka_topics

import (
	"fmt"
	"go.dfds.cloud/mcdonalds/selfservice"

	"github.com/spf13/cobra"
)

var (
	selfserviceClient *selfservice.SelfServiceClient
)

var TopicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "query topics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fetching topics")
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
		idInput, _ := cmd.Flags().GetString("id")
		fmt.Println(idInput)
		topics := selfserviceClient.GetTopics()
		fmt.Println(topics)
	},
}
