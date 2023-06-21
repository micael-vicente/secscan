package scan

type Config struct {
	ExcludedCheckers    []string
	SensitiveTerms      []string
	XSSTargetExtensions []string
}
