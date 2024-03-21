package capability

import (
	"fmt"
	"os"

	"go.dfds.cloud/ticli/selfservice"

	"github.com/spf13/cobra"
)

var (
	selfserviceClient *selfservice.SelfServiceClient
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
