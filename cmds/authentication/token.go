package authentication

import (
	"net/http"
	"os"
	"strings"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"go.dfds.cloud/ticli/cmds/configuration"
	"go.dfds.cloud/ticli/cmds/outputwriter"
)

func ResponseServer() {
	http.HandleFunc("/", serveFile)
	http.HandleFunc("/success", successHandler)
	http.ListenAndServe(":4200", nil)
}

func successHandler(resp http.ResponseWriter, req *http.Request) {
	box := rice.MustFindBox("web")
	buf, err := box.Bytes("success.html")
	if err != nil {
		outputwriter.GetWriter().WriteError(err)
	}

	err = req.ParseForm()
	if err != nil {
		outputwriter.GetWriter().WriteError(err)
	}

	accessToken := req.Form.Get("token")
	configuration.SaveAccessToken(accessToken)

	resp.WriteHeader(200)
	resp.Write(buf)

	go func() {
		time.Sleep(time.Second * 1)
		os.Exit(0)
	}()
}

func serveFile(resp http.ResponseWriter, req *http.Request) {
	box := rice.MustFindBox("web")

	httpBox := box.HTTPBox()
	path, _ := strings.CutPrefix(req.URL.Path, "/")

	// a hack
	if path == "login" {
		path = "login.html"
	}

	buf, err := httpBox.Bytes(path)
	if err != nil {
		bytes := []byte("Not found FLUTTERSHY")

		resp.WriteHeader(400)
		resp.Write(bytes)
		return
	}

	resp.WriteHeader(200)
	resp.Write(buf)
}
