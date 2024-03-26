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
	b := createCapabilityCmd
	CapabilityCmd.AddCommand(b)

	createCapabilityCmd.PersistentFlags().StringVar(&description, "description", "", "adds a description to a capability (required)")
	configuration.BindFlag("description", createCapabilityCmd.PersistentFlags().Lookup("description"))
	//createCapabilityCmd.MarkFlagRequired("description")
	createCapabilityCmd.PersistentFlags().StringVar(&metadata, "metadata", "", "add matadata JSON to a capability")
	configuration.BindFlag("metadata", createCapabilityCmd.PersistentFlags().Lookup("metadata"))

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
			Invitees:     []string{},
			JsonMetadata: metadata,
		}
		fmt.Println(capabilityStruct)
		capability := selfserviceClient.CreateCapability(capabilityStruct)
		fmt.Println(capability)

	},
}
