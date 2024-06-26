package authentication

import (
	"fmt"
	"net/url"

	"github.com/pkg/browser"
	"go.dfds.cloud/ticli/cmds/configuration"
	"go.dfds.cloud/ticli/cmds/outputwriter"

	"github.com/spf13/cobra"
)

var AuthenticationCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "Authenticate via Azure AD",
	Run: func(cmd *cobra.Command, args []string) {
		authenticateCmdFunction(cmd, args)
	},
}

func InitializeAuthentication() {
	AuthenticationCmd.AddCommand(clearInformationCmd)
	AuthenticationCmd.AddCommand(printTokenCmd)
}

func authenticateCmdFunction(cmd *cobra.Command, args []string) {
	const tenant = "73a99466-ad05-4221-9f90-e7142aa2f6c1"
	const scope = "api://3007f683-c3c2-4bf9-b6bd-2af03fb94f6d/.default"
	const appId = "3007f683-c3c2-4bf9-b6bd-2af03fb94f6d"

	urlParsed, err := url.Parse(fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/authorize", tenant))
	if err != nil {
		outputwriter.GetWriter().WriteError(err)
	}

	values := urlParsed.Query()
	values.Set("client_id", appId)
	values.Set("response_type", "token")
	values.Set("scope", scope)
	values.Set("nonce", "12345")
	values.Set("redirect_uri", "http://localhost:4200/login")
	urlParsed.RawQuery = values.Encode()

	browser.OpenURL(urlParsed.String())
	ResponseServer()
}

var clearInformationCmd = &cobra.Command{
	Use:   "clear",
	Short: "removes all stored authentication data",
	Run: func(cmd *cobra.Command, args []string) {
		configuration.ClearAccessToken()
	},
}

var printTokenCmd = &cobra.Command{
	Use:   "print",
	Short: "print the stored access token",
	Run: func(cmd *cobra.Command, args []string) {
		configuration.PrintAccessToken()
	},
}
