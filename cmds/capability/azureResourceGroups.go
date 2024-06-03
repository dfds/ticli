package capability

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
	"go.dfds.cloud/ticli/cmds/outputwriter"
	"go.dfds.cloud/ticli/openapiclient"
)

var AzureCmd = &cobra.Command{
	Use:   "azure",
	Short: "query azure",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func InitAzure(accessToken string) {
	AzureCmd.AddCommand(azureResourceGroupsCmd)
	azureResourceGroupsCmd.AddCommand(createResourceGroupCmd)
}

var azureResourceGroupsCmd = &cobra.Command{
	Use:   "resource-groups",
	Short: "get connected azure resource groups",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("id") {
			resourceGroups, err := selfserviceClient.GetCapabilitiesIdAzureresourcesWithResponse(context.Background(), CapabilityId)

			if err != nil {
				outputwriter.GetWriter().WriteError(err)
			}
			outputwriter.GetWriter().WriteData(resourceGroups.JSON200)
		} else {
			outputwriter.GetWriter().WriteError(NoIdError)
		}

	},
}

var createResourceGroupCmd = &cobra.Command{
	Use:   "create [ENVIRONMENT]",
	Short: "create a new azure resource group in the designated environment",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			environment := args[0]
			var resourceGroupRequest = openapiclient.NewAzureResourceRequest{
				Environment: environment,
			}

			resourceGroup, err := selfserviceClient.PostCapabilitiesIdAzureresourcesWithResponse(context.Background(), CapabilityId, resourceGroupRequest)
			if err != nil {
				outputwriter.GetWriter().WriteError(err)
			}
			outputwriter.GetWriter().WriteData(resourceGroup.JSON200)
		} else {
			outputwriter.GetWriter().WriteError(errors.New("No environment specified"))
		}
	},
}
