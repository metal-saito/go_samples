package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-samples",
	Short: "Go language samples for reservation tool",
	Long:  "This is a collection of Go language samples demonstrating CLI, HTTP API, and worker pool patterns.",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

