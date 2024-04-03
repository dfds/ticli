package capability

import (
	"fmt"
	"go.dfds.cloud/ticli/cmds/configuration"
	"os"

	"go.dfds.cloud/ticli/selfservice"

	"github.com/spf13/cobra"
)

var (
	selfserviceClient *selfservice.SelfServiceClient

	description string
	metadata    string
	invitees    []string
)

var CapabilityCmd = &cobra.Command{
	Use:   "capability",
	Short: "query capabilities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Can I take your order")
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

	selfserviceClient = selfservice.NewSelfServiceClient(accessToken)
}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query the API",
	Run: func(cmd *cobra.Command, args []string) {
		capabilities := selfserviceClient.GetCapabilities()
		fmt.Println(capabilities)
	},
}

var capabilityByIdCmd = &cobra.Command{
	Use:   "id [ID]",
	Short: "returns a capability by id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Missing id")
			os.Exit(1)
		}
		capabilityById := selfserviceClient.GetCapabilityByID(args[0])
		fmt.Println(capabilityById)

	},
}

var createCapabilityCmd = &cobra.Command{
	Use:   "create [NAME]",
	Short: "creates a new capability",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Missing id")
			os.Exit(1)
		}
		capabilityStruct := selfservice.CapabilityRequest{
			Name:         args[0],
			Description:  description,
			Invitees:     invitees,
			JsonMetadata: metadata,
		}

		capability := selfserviceClient.CreateCapability(capabilityStruct)
		fmt.Println(capability)

	},
}
