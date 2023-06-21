package report

type Writer interface {
	WriteReport(results []Result, config Config)
}
