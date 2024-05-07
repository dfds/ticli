package capability

import (
	"context"
	"fmt"
	"go.dfds.cloud/ticli/openapiclient"
	"log"
	"os"

	"go.dfds.cloud/ticli/cmds/configuration"

	"go.dfds.cloud/ticli/cmds/outputwriter"
	"go.dfds.cloud/ticli/selfservice"

	"github.com/spf13/cobra"
)

var (
	selfserviceClient *openapiclient.ClientWithResponses

	description string
	metadata    string
	invitees    []string
)

var CapabilityCmd = &cobra.Command{
	Use:   "capability",
	Short: "query capabilities",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func InitializeCapability(accessToken string) {
	CapabilityCmd.AddCommand(queryCmd)
	CapabilityCmd.AddCommand(capabilityByIdCmd)
	CapabilityCmd.AddCommand(createCapabilityCmd)

	createCapabilityCmd.PersistentFlags().StringVar(&description, "description", "", "adds a description to a capability (required)")
	configuration.BindFlag("description", createCapabilityCmd.PersistentFlags().Lookup("description"))
	cobra.MarkFlagRequired(createCapabilityCmd.PersistentFlags(), "description")
	createCapabilityCmd.PersistentFlags().StringVar(&metadata, "metadata", "", "add matadata JSON to a capability; the only required field of the JSON is the cost centre (required)")
	configuration.BindFlag("metadata", createCapabilityCmd.PersistentFlags().Lookup("metadata"))
	cobra.MarkFlagRequired(createCapabilityCmd.PersistentFlags(), "metadata")
	createCapabilityCmd.PersistentFlags().StringArrayVar(&invitees, "invitees", []string{}, "add invitees array to a capability")
	configuration.BindFlag("invitees", createCapabilityCmd.PersistentFlags().Lookup("invitees"))

	selfserviceClient = selfservice.NewGeneratedClient(accessToken)
}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query the API",
	Run: func(cmd *cobra.Command, args []string) {

		capabilities, err := selfserviceClient.GetCapabilitiesWithResponse(context.Background())

		if err != nil {
			log.Fatal(err)
		}

		outputwriter.GetWriter().WriteData(capabilities.JSON200)
	},
}

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
			log.Fatal(err)
		}

		outputwriter.GetWriter().WriteData(capability.JSON200)

	},
}

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

		fmt.Println("before post")

		capability, err := selfserviceClient.PostCapabilitiesWithResponse(context.Background(), capabilityStruct)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("after post")
		outputwriter.GetWriter().WriteData(capability.JSON201)

	},
}
