package capability

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"go.dfds.cloud/ticli/openapiclient"

	"go.dfds.cloud/ticli/cmds/configuration"

	"go.dfds.cloud/ticli/cmds/outputwriter"
	"go.dfds.cloud/ticli/selfservice"

	"github.com/spf13/cobra"
)

var (
	selfserviceClient *openapiclient.ClientWithResponses

	description         string
	metadata            string
	invitees            []string
	CapabilityId        string
	RecommendationLevel string

	NoIdError = errors.New("No id specified")
)

var CapabilityCmd = &cobra.Command{
	Use:   "capability", // Todo: change to "capabilities"
	Short: "query capabilities",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func InitializeCapability(accessToken string) {
	CapabilityCmd.AddCommand(queryCmd)
	//CapabilityCmd.AddCommand(capabilityByIdCmd)
	CapabilityCmd.AddCommand(createCapabilityCmd)
	CapabilityCmd.AddCommand(recommendationsCmd)

	CapabilityCmd.PersistentFlags().StringVar(&CapabilityId, "id", "", "define capability id")

	createCapabilityCmd.PersistentFlags().StringVar(&description, "description", "", "adds a description to a capability (required)")
	configuration.BindFlag("description", createCapabilityCmd.PersistentFlags().Lookup("description"))
	cobra.MarkFlagRequired(createCapabilityCmd.PersistentFlags(), "description")
	createCapabilityCmd.PersistentFlags().StringVar(&metadata, "metadata", "", "add matadata JSON to a capability; the only required field of the JSON is the cost centre (required)")
	configuration.BindFlag("metadata", createCapabilityCmd.PersistentFlags().Lookup("metadata"))
	cobra.MarkFlagRequired(createCapabilityCmd.PersistentFlags(), "metadata")
	createCapabilityCmd.PersistentFlags().StringArrayVar(&invitees, "invitees", []string{}, "add invitees array to a capability")
	configuration.BindFlag("invitees", createCapabilityCmd.PersistentFlags().Lookup("invitees"))

	recommendationsCmd.PersistentFlags().StringVar(&RecommendationLevel, "level", "", "check desired recommendation level (partial, full, none)")

	selfserviceClient = selfservice.NewGeneratedClient(accessToken)

	InitAWS(configuration.GetString("access-token"))
	CapabilityCmd.AddCommand(AWSCmd)

	InitAzure(configuration.GetString("access-token"))
	CapabilityCmd.AddCommand(AzureCmd)

}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query the API",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("id") { // Id is set; Get specific capability
			capability, err := selfserviceClient.GetCapabilitiesIdWithResponse(context.Background(), CapabilityId)
			if err != nil {
				outputwriter.GetWriter().WriteError(err)
			}
			outputwriter.GetWriter().WriteData(capability.JSON200)
		} else { // Id is not set; Get all capabilities
			capabilities, err := selfserviceClient.GetCapabilitiesWithResponse(context.Background())
			if err != nil {
				outputwriter.GetWriter().WriteError(err)
			}
			outputwriter.GetWriter().WriteData(capabilities.JSON200)
		}
	},
}

/*
var capabilityByIdCmd = &cobra.Command{
	Use:   "id [ID]",
	Short: "returns a capability by id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			outputwriter.GetWriter().WriteError(fmt.Errorf("missing id"))
			os.Exit(1)
		}

		capability, err := selfserviceClient.GetCapabilitiesIdWithResponse(context.Background(), args[0])

		if err != nil {
			outputwriter.GetWriter().WriteError(err)
		}

		outputwriter.GetWriter().WriteData(capability.JSON200)

	},
}
*/

var createCapabilityCmd = &cobra.Command{
	Use:   "create [NAME]",
	Short: "creates a new capability",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			outputwriter.GetWriter().WriteError(fmt.Errorf("missing name"))
			os.Exit(1)
		}
		capabilityStruct := openapiclient.NewCapabilityRequest{
			Name:         args[0],
			Description:  &description,
			Invitees:     &invitees,
			JsonMetadata: &metadata,
		}

		capability, err := selfserviceClient.PostCapabilitiesWithResponse(context.Background(), capabilityStruct)

		if err != nil {
			outputwriter.GetWriter().WriteError(err)
		}
		outputwriter.GetWriter().WriteData(capability.JSON201)

	},
}

var recommendationsCmd = &cobra.Command{
	Use:   "recommendations",
	Short: "list recommendations",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("id") { // Id is set; Get specific capability
			configurationLevel, err := selfserviceClient.GetCapabilitiesIdConfigurationlevelWithResponse(context.Background(), CapabilityId)
			if err != nil {
				outputwriter.GetWriter().WriteError(err)
			}

			var recommendation map[string]interface{}
			err = json.Unmarshal(configurationLevel.Body, &recommendation)
			if err != nil {
				outputwriter.GetWriter().WriteError(errors.New(fmt.Sprint("Error parsing JSON:", err)))
			}

			if cmd.Flags().Changed("level") { // desired level set, return goodly if it matches
				if recommendation["overallLevel"] == RecommendationLevel {
					outputwriter.GetWriter().WriteData("Capability configuration level matches expectations")
				} else {
					outputwriter.GetWriter().WriteError(errors.New("capability configuration level does not match expectations"))
				}
			} else { // desired level not set, return all recommendations
				outputwriter.GetWriter().WriteData(recommendation)
			}
		} else {
			outputwriter.GetWriter().WriteError(NoIdError)
		}
	},
}
