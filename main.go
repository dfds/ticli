package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.dfds.cloud/ticli/cmds"
)

var (
	version = "0.0"
	repositoryUrl = "https://api.github.com/repos/dfds/ticli/tags"
)

type Tag struct {
	Name string `json:"name"`
}

func init() {
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
		fmt.Println("Latest tag:", latestTag)
		if latestTag != version {
			panic("Version Check: New version of Ticli available, please download the newest version")
		}
	} else {
		fmt.Println("Version Check: No version found")
	}
}

func main() {
	cmds.Execute()
}
