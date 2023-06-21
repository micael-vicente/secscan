package secscancli

import (
	"ACME/secscan/pkg/report"
	"ACME/secscan/pkg/scan"
	"ACME/secscan/pkg/sensors"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var scanCmd = &cobra.Command{
	Use: "scan [flags] [directory]",
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		produceJson, _ := cmd.Flags().GetBool("json")
		Scan(args, configPath, produceJson)
	},
}

// Scan runs the security scan for a given command
func Scan(args []string, configFile string, wantsJsonOutput bool) {

	if len(args) < 1 {
		log.Fatalf("directory not provided")
	}

	var config Config
	if _, err := os.Stat(configFile); err == nil {
		config = ReadConfig(configFile)
	} else {
		config = ReadConfig("./config.yml")
	}

	files := sensors.RunDirectory(args[0], GetSensorsConfig(config))

	log.Printf("sensors found %d files to scan", len(files))

	var findings []scan.Finding
	for _, f := range files {
		log.Printf("scanning file '%s'", f)
		file, _ := os.ReadFile(f)
		findings = append(findings, scan.Scan(f, file, GetScanConfig(config))...)
	}

	reportConfig := GetReportConfig(config)
	if wantsJsonOutput {
		reportConfig.Format = "json"
	}
	report.WriteReport(convertFindings(findings), reportConfig)
}

func convertFindings(findings []scan.Finding) []report.Result {
	var results []report.Result

	for _, f := range findings {
		results = append(results, report.Result{Code: f.Code, Line: f.Line, File: f.File})
	}

	return results
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().BoolP("json", "j", false, "output in json format")
	scanCmd.Flags().StringP("config", "c", "", "path and name of config file, ex: ./config.yml")
	_ = rootCmd.MarkFlagRequired("scan")
}
