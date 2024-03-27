package outputwriter

import "fmt"

type YamlWriter struct{}

func (y *YamlWriter) WriteData(data interface{}) {
	fmt.Println("YamlWriter Data: NOTHING!!!")
}

func (y *YamlWriter) WriteError(err error) {
	fmt.Println("YamlWriter Error: ", err)
}

func CreateYamlWriter() OutputWriter {
	return &YamlWriter{}
}
