package outputwriter

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonError struct {
	Error string `json:"error"`
}

type JsonWriter struct{}

func (j *JsonWriter) WriteData(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		j.WriteError(err)
		return
	}
	// we should not be in charge of printing, only formating
	fmt.Println(string(jsonData))
}

func (j *JsonWriter) WriteError(err error) {
	e := JsonError{Error: err.Error()}
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		j.WriteError(err)
		os.Exit(1) // errors should indicate issues also via exit code
	}
	fmt.Println(string(jsonData))
}
func CreateJsonWriter() OutputWriter {
	return &JsonWriter{}
}
