package secscancli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "secscan scan [flags] [directory]",
	Short: "You know, to help you secure your code",
}

// Execute to call the root cmd
func Execute() error {
	return rootCmd.Execute()
}
