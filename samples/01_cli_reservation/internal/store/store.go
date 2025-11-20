package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reservation_tool/internal/domain"
)

// Store は予約データの永続化を担当します
type Store struct {
	filePath string
}

// NewStore は新しいStoreを作成します
func NewStore(filePath string) *Store {
	return &Store{filePath: filePath}
}

// Add は予約を追加します
func (s *Store) Add(reservation *domain.Reservation) error {
	reservations, err := s.load()
	if err != nil {
		return err
	}

	// 重複チェック
	for _, r := range reservations {
		if r.Status == "booked" && reservation.Overlaps(r) {
			return fmt.Errorf("時間が重複している予約が既に存在します")
		}
	}

	reservations = append(reservations, reservation)
	return s.save(reservations)
}

// All はすべての予約を取得します
func (s *Store) All() ([]*domain.Reservation, error) {
	return s.load()
}

// Cancel は予約をキャンセルします
func (s *Store) Cancel(id string) error {
	reservations, err := s.load()
	if err != nil {
		return err
	}

	found := false
	for _, r := range reservations {
		if r.ID == id {
			r.Status = "cancelled"
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("指定IDの予約は存在しません: %s", id)
	}

	return s.save(reservations)
}

func (s *Store) load() ([]*domain.Reservation, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []*domain.Reservation{}, nil
		}
		return nil, err
	}

	var reservations []*domain.Reservation
	if len(data) == 0 {
		return []*domain.Reservation{}, nil
	}

	if err := json.Unmarshal(data, &reservations); err != nil {
		return nil, fmt.Errorf("JSON解析エラー: %w", err)
	}

	return reservations, nil
}

func (s *Store) save(reservations []*domain.Reservation) error {
	// ディレクトリが存在しない場合は作成
	dir := filepath.Dir(s.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("ディレクトリ作成エラー: %w", err)
	}

	data, err := json.MarshalIndent(reservations, "", "  ")
	if err != nil {
		return fmt.Errorf("JSONシリアライズエラー: %w", err)
	}

	if err := os.WriteFile(s.filePath, data, 0644); err != nil {
		return fmt.Errorf("ファイル書き込みエラー: %w", err)
	}

	return nil
}

