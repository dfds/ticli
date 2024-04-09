package outputwriter

func CreateDefaultWriter() OutputWriter {
	return &JsonWriter{}
}
