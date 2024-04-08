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

func SetWriter(w OutputWriter) {
	writer = w
}
