package scan

import (
	"log"
	"regexp"
	"strings"
)

const SQLIErrorCode = "[SQL Injection]"

type SQLInjectionChecker struct {
}

func (c SQLInjectionChecker) Perform(fileName string, fileContent []byte, config Config) []Finding {
	// find ".... SELECT .... WHERE .... %s .... "
	log.Println("running SQL injection checks...")

	queryPattern := "(?i)\"+[\\s\\S]*(select)+[\\s\\S]*(where)+[\\s\\S]*(%s)+[\\s\\S]*\"+"
	match, _ := regexp.Match(queryPattern, fileContent)
	var findings []Finding

	// no need to verify file if it does not contain a query with requirements
	if !match {
		return findings
	}

	strContent := string(fileContent)
	lines := strings.Split(strContent, "\n")

	r, _ := regexp.Compile(queryPattern)

	for i, l := range lines {
		if s := r.FindString(l); s != "" {
			findings = append(findings, Finding{File: fileName, Code: SQLIErrorCode, Line: i + 1})
		}
	}

	return findings
}

func (c SQLInjectionChecker) GetIdentifier() string {
	return "SQLInjectionChecker"
}
