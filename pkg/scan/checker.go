package scan

// Checker checker type
type Checker interface {
	// Perform performs analysis to the given file content, produces a scan.Result.
	Perform(fileName string, fileContent []byte, config Config) []Finding
	// GetIdentifier returns the identifier of a Checker type
	GetIdentifier() string
}
