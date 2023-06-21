package scan

import (
	"log"
	"strings"
)

const SDEErrorCode = "[Sensitive Data Exposure]"

type SDEChecker struct {
}

func (c SDEChecker) Perform(fileName string, fileContent []byte, config Config) []Finding {
	log.Println("running sensitive data exposure checks...")
	var findings []Finding

	if len(config.SensitiveTerms) == 0 {
		return findings
	}

	strContent := string(fileContent)
	lines := strings.Split(strContent, "\n")

	termsFound := 0

	for i, l := range lines {
		for _, s := range config.SensitiveTerms {
			contains := strings.Contains(l, s)
			if !contains {
				break
			}
			termsFound++
		}

		if termsFound == len(config.SensitiveTerms) {
			findings = append(findings, Finding{File: fileName, Line: i + 1, Code: SDEErrorCode})
		}
		termsFound = 0
	}

	return findings

}

func (c SDEChecker) GetIdentifier() string {
	return "SDEChecker"
}
