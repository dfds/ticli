package authentication

import (
	"fmt"
	"net/url"

	"github.com/pkg/browser"
	log "github.com/sirupsen/logrus"

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
	// nothing to see here yet, but the wiring now exists
}

func authenticateCmdFunction(cmd *cobra.Command, args []string) {
	const tenant = "73a99466-ad05-4221-9f90-e7142aa2f6c1"
	const scope = "api://3007f683-c3c2-4bf9-b6bd-2af03fb94f6d/.default"
	const appId = "3007f683-c3c2-4bf9-b6bd-2af03fb94f6d"

	urlParsed, err := url.Parse(fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/authorize", tenant))
	if err != nil {
		log.Fatal(err)
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
