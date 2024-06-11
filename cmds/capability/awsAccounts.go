package capability

import (
	"context"

	"go.dfds.cloud/ticli/cmds/outputwriter"

	"github.com/spf13/cobra"
)

var AWSCmd = &cobra.Command{
	Use:   "aws",
	Short: "query topics",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func InitAWS(accessToken string) {
	AWSCmd.AddCommand(getAccountCmd)
	getAccountCmd.AddCommand(createAccountCmd)
}

var getAccountCmd = &cobra.Command{
	Use:   "account",
	Short: "get connected aws account",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("id") {
			topics, err := selfserviceClient.GetCapabilitiesIdAwsaccountWithResponse(context.Background(), CapabilityId)
			if err != nil {
				outputwriter.GetWriter().WriteError(err)
			}
			outputwriter.GetWriter().WriteData(topics.JSON200)
		} else {
			outputwriter.GetWriter().WriteError(NoIdError)
		}
	},
}

var createAccountCmd = &cobra.Command{
	Use:   "create",
	Short: "create an aws account for this capability",
	Run: func(cmd *cobra.Command, args []string) {
		account, err := selfserviceClient.PostCapabilitiesIdAwsaccountWithResponse(context.Background(), CapabilityId)
		if err != nil {
			outputwriter.GetWriter().WriteError(err)
		}
		outputwriter.GetWriter().WriteData(account.JSON200)
	},
}
