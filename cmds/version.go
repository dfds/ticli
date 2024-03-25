package cmds

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		fmt.Println("Version Check: Error fetching Ticli tags:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var tags []Tag
	err = json.NewDecoder(resp.Body).Decode(&tags)
	if err != nil {
		fmt.Println("Version Check: Error parsing JSON response:", err)
		return
	}

	// Assuming tags are returned in chronological order, the first tag will be the latest
	if len(tags) > 0 {
		latestTag := tags[0].Name
		if latestTag != Version {
			panic("Version Check: New version of Ticli available, please update to the latest version")
		}
	} else {
		fmt.Println("Version Check: No remote version found, ignoring version verification")
	}
}
