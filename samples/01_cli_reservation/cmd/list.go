package cmd

import (
	"fmt"
	"sort"
	"time"

	"github.com/spf13/cobra"
	"reservation_tool/internal/store"
)

// ListCmd は予約一覧表示コマンドです
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "予約一覧を表示",
	Long:  "登録されている予約の一覧を表示します。",
	RunE: func(c *cobra.Command, args []string) error {
		s := store.NewStore("data/reservations.json")
		reservations, err := s.All()
		if err != nil {
			return fmt.Errorf("予約の取得エラー: %w", err)
		}

		if len(reservations) == 0 {
			fmt.Println("予約はまだありません。")
			return nil
		}

		sort.Slice(reservations, func(i, j int) bool {
			return reservations[i].StartsAt.Before(reservations[j].StartsAt)
		})

		fmt.Printf("%-8s | %-6s | %-8s | %-24s | %-24s\n", "ID", "利用者", "リソース", "開始", "終了")
		for _, r := range reservations {
			fmt.Printf("%-8s | %-6s | %-8s | %-24s | %-24s\n",
				r.ID, r.UserName, r.ResourceName, r.StartsAt.Format(time.RFC3339), r.EndsAt.Format(time.RFC3339))
		}

		return nil
	},
}

func init() {
	// main.go の rootCmd に追加される
}

