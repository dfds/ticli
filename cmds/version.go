package cmds

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.dfds.cloud/ticli/cmds/outputwriter"
)

var (
	Version       = "0.0"
	repositoryUrl = "https://api.github.com/repos/dfds/ticli/tags"
)

type Tag struct {
	Name string `json:"name"`
}

func remoteVersionCheck() {
	// Send GET request to GitHub API
	resp, err := http.Get(repositoryUrl)
	if err != nil {
		outputwriter.GetWriter().WriteError(fmt.Errorf("version check: error while getting ticli tags: %s", err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var tags []Tag
	err = json.NewDecoder(resp.Body).Decode(&tags)
	if err != nil {
		outputwriter.GetWriter().WriteError(fmt.Errorf("version check: error while parsing JSON response: %s", err))
		os.Exit(1)
	}

	// Assuming tags are returned in chronological order, the first tag will be the latest
	if len(tags) > 0 {
		latestTag := tags[0].Name
		if latestTag != Version {
			outputwriter.GetWriter().WriteError(fmt.Errorf("version Check: New version of Ticli available, please update to the latest version"))
			os.Exit(1)
		}
	}
}
