package report

import "log"

var lookup = map[string]Writer{
	// add new writers here
	"json": JsonWriter{},
	"text": TextWriter{},
}

func WriteReport(results []Result, config Config) {
	writer := lookup[config.Format]

	if writer == nil {
		log.Fatalf("ResultWriter for type: [%s] not registered", config.Format)
	}

	writer.WriteReport(results, config)
}
