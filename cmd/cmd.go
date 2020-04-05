package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:     "main",
	Short:   "Go Restful API Boilerplate",
	Version: "1.0.0-alpha",
	Example: "main",
}

// Execute main external command
func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
