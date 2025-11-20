package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"reservation_tool/internal/domain"
	"reservation_tool/internal/store"
)

// AddCmd は予約追加コマンドです
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "予約を追加",
	Long:  "新しい予約を登録します。",
	RunE: func(c *cobra.Command, args []string) error {
		name, _ := c.Flags().GetString("name")
		resource, _ := c.Flags().GetString("resource")
		startsAtStr, _ := c.Flags().GetString("starts-at")
		endsAtStr, _ := c.Flags().GetString("ends-at")

		if name == "" || resource == "" || startsAtStr == "" || endsAtStr == "" {
			return fmt.Errorf("必須パラメータが不足しています")
		}

		startsAt, err := time.Parse(time.RFC3339, startsAtStr)
		if err != nil {
			return fmt.Errorf("開始時刻の解析エラー: %w", err)
		}

		endsAt, err := time.Parse(time.RFC3339, endsAtStr)
		if err != nil {
			return fmt.Errorf("終了時刻の解析エラー: %w", err)
		}

		reservation, err := domain.NewReservation(name, resource, startsAt, endsAt)
		if err != nil {
			return fmt.Errorf("予約の作成エラー: %w", err)
		}

		s := store.NewStore("data/reservations.json")
		if err := s.Add(reservation); err != nil {
			return fmt.Errorf("予約の保存エラー: %w", err)
		}

		fmt.Printf("予約を登録しました: %s\n", reservation.ID)
		return nil
	},
}

func init() {
	// main.go の rootCmd に追加される
}
func init() {
	AddCmd.Flags().String("name", "", "利用者名 (必須)")
	AddCmd.Flags().String("resource", "", "リソース名 (必須)")
	AddCmd.Flags().String("starts-at", "", "開始時刻 (ISO8601形式, 必須)")
	AddCmd.Flags().String("ends-at", "", "終了時刻 (ISO8601形式, 必須)")
	AddCmd.MarkFlagRequired("name")
	AddCmd.MarkFlagRequired("resource")
	AddCmd.MarkFlagRequired("starts-at")
	AddCmd.MarkFlagRequired("ends-at")
}
}

