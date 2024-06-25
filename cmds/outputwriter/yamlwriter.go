package outputwriter

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type YamlError struct {
	Error string
}

type YamlWriter struct{}

func (y *YamlWriter) WriteData(data interface{}) {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		y.WriteError(err)
		return
	}
	// we should not be in charge of printing, only formating
	fmt.Println(string(yamlData))
}

func (y *YamlWriter) WriteError(err error) {
	e := YamlError{Error: err.Error()}
	yamlData, err := yaml.Marshal(e)
	if err != nil {
		y.WriteError(err)
		os.Exit(1) // errors should indicate issues also via exit code
	}
	fmt.Println(string(yamlData))
}
func CreateYamlWriter() OutputWriter {
	return &YamlWriter{}
}
