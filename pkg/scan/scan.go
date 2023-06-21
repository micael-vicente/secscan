package scan

import "ACME/secscan/pkg/util"

var checkers = []Checker{
	SQLInjectionChecker{},
	XSSChecker{},
	SDEChecker{},
}

// Scan scans given file using registered checkers that have not been excluded
func Scan(fileName string, fileContent []byte, config Config) []Finding {

	var findings []Finding

	for _, c := range checkers {
		skipChecker := util.Contains(config.ExcludedCheckers, c.GetIdentifier())

		if !skipChecker {
			findings = append(findings, c.Perform(fileName, fileContent, config)...)
		}
	}

	return findings
}
