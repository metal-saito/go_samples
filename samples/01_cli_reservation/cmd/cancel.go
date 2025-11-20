package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"reservation_tool/internal/store"
)

// CancelCmd は予約キャンセルコマンドです
var CancelCmd = &cobra.Command{
	Use:   "cancel [ID]",
	Short: "予約をキャンセル",
	Long:  "指定されたIDの予約をキャンセルします。",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		id := args[0]
		s := store.NewStore("data/reservations.json")

		if err := s.Cancel(id); err != nil {
			return fmt.Errorf("予約のキャンセルエラー: %w", err)
		}

		fmt.Printf("予約をキャンセルしました: %s\n", id)
		return nil
	},
}

func init() {
	// main.go の rootCmd に追加される
}

