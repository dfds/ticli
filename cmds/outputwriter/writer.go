package outputwriter

var (
	writer OutputWriter
)

type OutputWriter interface {
	WriteData(data interface{})
	WriteError(err error)
}

func GetWriter() OutputWriter {
	return writer
}

func SetWriter(writerType string) {
	switch writerType {
	case "json":
		writer = CreateJsonWriter()
	case "yaml":
		writer = CreateYamlWriter()
	default:
		writer = CreateJsonWriter()
	}
}
