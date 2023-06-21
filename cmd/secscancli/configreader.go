package secscancli

import (
	"ACME/secscan/pkg/report"
	"ACME/secscan/pkg/scan"
	"ACME/secscan/pkg/sensors"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var config Config

type Config struct {
	Scan struct {
		ExcludedCheckers    []string `yaml:"excludedCheckers"`
		SensitiveTerms      []string `yaml:"sensitiveTerms"`
		XSSTargetExtensions []string `yaml:"xSSTargetExtensions"`
	} `yaml:"scan"`
	Report struct {
		OutputPath string `yaml:"outputPath"`
		Format     string `yaml:"format"`
	} `yaml:"report"`
	Sensors struct {
		ExcludedFileExtensions []string `yaml:"excludedFileExtensions"`
	} `yaml:"sensors"`
}

// ReadConfig given a file path, unmarshalls it into a struct
func ReadConfig(file string) Config {
	f, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Impossible to open config file '%s'", file)
	}

	err = yaml.Unmarshal(f, &config)

	if err != nil {
		log.Fatalf("Impossible unmarshal config file '%s'", file)
	}

	return config
}

// GetScanConfig gets the scan config from the general config struct
func GetScanConfig(config Config) scan.Config {
	return scan.Config{
		ExcludedCheckers:    config.Scan.ExcludedCheckers,
		SensitiveTerms:      config.Scan.SensitiveTerms,
		XSSTargetExtensions: config.Scan.XSSTargetExtensions,
	}
}

// GetSensorsConfig gets the sensor config from the general config struct
func GetSensorsConfig(config Config) sensors.Config {
	return sensors.Config{
		ExcludedFileExtensions: config.Sensors.ExcludedFileExtensions,
	}
}

// GetReportConfig gets the report config from the general config struct
func GetReportConfig(config Config) report.Config {
	return report.Config{
		Format:     config.Report.Format,
		OutputPath: config.Report.OutputPath,
	}
}
