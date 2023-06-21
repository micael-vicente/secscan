package report

import (
	"encoding/json"
	"os"
)

type JsonWriter struct {
}

func (jw JsonWriter) WriteReport(results []Result, config Config) {
	_ = os.RemoveAll(config.OutputPath + "/result.json")
	f, _ := os.Create(config.OutputPath + "/result.json")

	content, _ := json.MarshalIndent(results, "", "  ")

	_, _ = f.Write(content)
}
