package store

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"reservation_tool/internal/domain"
)

func TestStore_Add(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "test.json")
	s := NewStore(tmpFile)

	reservation, err := domain.NewReservation("Alice", "Room-A", time.Now(), time.Now().Add(time.Hour))
	if err != nil {
		t.Fatalf("予約の作成に失敗: %v", err)
	}

	if err := s.Add(reservation); err != nil {
		t.Fatalf("予約の追加に失敗: %v", err)
	}

	all, err := s.All()
	if err != nil {
		t.Fatalf("予約の取得に失敗: %v", err)
	}

	if len(all) != 1 {
		t.Fatalf("予約数が期待と異なります: 期待=1, 実際=%d", len(all))
	}

	if all[0].ID != reservation.ID {
		t.Errorf("予約IDが期待と異なります: 期待=%s, 実際=%s", reservation.ID, all[0].ID)
	}
}

func TestStore_Cancel(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "test.json")
	s := NewStore(tmpFile)

	reservation, err := domain.NewReservation("Alice", "Room-A", time.Now(), time.Now().Add(time.Hour))
	if err != nil {
		t.Fatalf("予約の作成に失敗: %v", err)
	}

	if err := s.Add(reservation); err != nil {
		t.Fatalf("予約の追加に失敗: %v", err)
	}

	if err := s.Cancel(reservation.ID); err != nil {
		t.Fatalf("予約のキャンセルに失敗: %v", err)
	}

	all, err := s.All()
	if err != nil {
		t.Fatalf("予約の取得に失敗: %v", err)
	}

	if all[0].Status != "cancelled" {
		t.Errorf("予約ステータスが期待と異なります: 期待=cancelled, 実際=%s", all[0].Status)
	}
}

func TestStore_OverlapCheck(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "test.json")
	s := NewStore(tmpFile)

	baseTime := time.Now()
	r1, _ := domain.NewReservation("Alice", "Room-A", baseTime, baseTime.Add(time.Hour))
	r2, _ := domain.NewReservation("Bob", "Room-A", baseTime.Add(30*time.Minute), baseTime.Add(90*time.Minute))

	if err := s.Add(r1); err != nil {
		t.Fatalf("予約1の追加に失敗: %v", err)
	}

	if err := s.Add(r2); err == nil {
		t.Error("重複する予約が追加されてしまいました")
	}
}

