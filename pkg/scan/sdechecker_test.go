package scan

import "testing"

// given single term, single line, line contains term, then return 1 finding
func TestSDEChecker_PerformSingleTermFound(t *testing.T) {

	content := []byte("<p>ACME acquired by Tom & Jerry. $1.15b was the value of the acquisition</p>")

	findings := SDEChecker{}.Perform("myfile", content, Config{
		SensitiveTerms: []string{"ACME"},
	})

	actual := len(findings)
	expected := 1
	if actual != expected {
		t.Fatalf("expected %d findings but got %d", expected, actual)
	}

}

// given multiple terms, single line, line contains all terms, then return 1 finding
func TestSDEChecker_PerformMultipleTermsFound(t *testing.T) {

	content := []byte("<p>ACME acquired by Tom & Jerry. $1.15b was the value of the acquisition</p>")

	findings := SDEChecker{}.Perform("myfile", content, Config{
		SensitiveTerms: []string{"ACME", "Tom"},
	})

	actual := len(findings)
	expected := 1
	if actual != expected {
		t.Fatalf("expected %d findings but got %d", expected, actual)
	}

}

// given no terms, single line, line should be ok
func TestSDEChecker_PerformNoTermsLineIsFine(t *testing.T) {

	content := []byte("<p>ACME acquired by Tom & Jerry. $1.15b was the value of the acquisition</p>")

	findings := SDEChecker{}.Perform("myfile", content, Config{
		SensitiveTerms: []string{},
	})

	actual := len(findings)
	expected := 0
	if actual != expected {
		t.Fatalf("expected %d findings but got %d", expected, actual)
	}

}

// given multiple terms, multiple lines, some terms found, then no findings
func TestSDEChecker_PerformSomeTermsFound(t *testing.T) {

	content := []byte("<p>ACME acquired by Tom & Jerry. $1.15b was the value of the acquisition</p>\n" +
		"<p>ACME acquired by Tom & Jerry. $1.15b was the value of the acquisition</p>\n" +
		"<p>ACME acquired by Tom & Jerry. $1.15b was the value of the acquisition</p>")

	findings := SDEChecker{}.Perform("myfile", content, Config{
		SensitiveTerms: []string{"Potatoes", "ACME"},
	})

	actual := len(findings)
	expected := 0
	if actual != expected {
		t.Fatalf("expected %d findings but got %d", expected, actual)
	}

}

// given multiple terms, multiple lines, no terms found, then ok
func TestSDEChecker_PerformNoTermsFound(t *testing.T) {

	content := []byte("<p>ACME acquired by Tom & Jerry. $1.15b was the value of the acquisition</p>\n" +
		"<p>ACME acquired by Tom & Jerry. $1.15b was the value of the acquisition</p>\n" +
		"<p>ACME acquired by Tom & Jerry. $1.15b was the value of the acquisition</p>")

	findings := SDEChecker{}.Perform("myfile", content, Config{
		SensitiveTerms: []string{"Potatoes", "Tomatoes"},
	})

	actual := len(findings)
	expected := 0
	if actual != expected {
		t.Fatalf("expected %d findings but got %d", expected, actual)
	}

}
