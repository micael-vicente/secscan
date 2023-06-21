package scan

import (
	"ACME/secscan/pkg/util"
	"log"
	"path"
	"regexp"
	"strings"
)

const XSSErrorCode = "[Cross Site Scripting]"

type XSSChecker struct {
}

func (c XSSChecker) Perform(fileName string, fileContent []byte, config Config) []Finding {
	//find Alert() and probably Alert("something")
	ext := path.Ext(fileName)
	if !util.Contains(config.XSSTargetExtensions, ext) {
		return []Finding{}
	}

	log.Println("running sensitive data exposure checks...")

	pattern := "[aA]lert\\("
	r, _ := regexp.Compile(pattern)

	strContent := string(fileContent)
	lines := strings.Split(strContent, "\n")
	var findings []Finding

	for i, l := range lines {
		if s := r.FindString(l); s != "" {
			findings = append(findings, Finding{File: fileName, Line: i + 1, Code: XSSErrorCode})
		}
	}

	return findings
}

func (c XSSChecker) GetIdentifier() string {
	return "XSSChecker"
}
