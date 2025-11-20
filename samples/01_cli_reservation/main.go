package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"reservation_tool/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "reserve",
	Short: "予約管理CLIツール",
	Long:  "予約の登録、一覧表示、キャンセルを行うCLIツールです。",
}

func main() {
	rootCmd.AddCommand(cmd.AddCmd)
	rootCmd.AddCommand(cmd.ListCmd)
	rootCmd.AddCommand(cmd.CancelCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

