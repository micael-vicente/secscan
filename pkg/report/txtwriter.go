package report

import (
	"fmt"
	"os"
)

type TextWriter struct {
}

func (t TextWriter) WriteReport(results []Result, config Config) {

	_ = os.RemoveAll(config.OutputPath + "/result.txt")
	f, _ := os.Create(config.OutputPath + "/result.txt")

	for _, r := range results {
		format := "%s in file “%s” on line %d\n"
		_, _ = f.WriteString(fmt.Sprintf(format, r.Code, r.File, r.Line))
	}
}
