package repository

import (
	"fmt"
	"sync"

	"reservation_api/internal/domain"
)

// MemoryRepository はメモリ上で予約を管理するリポジトリです
type MemoryRepository struct {
	mu           sync.RWMutex
	reservations map[string]*domain.Reservation
}

// NewMemoryRepository は新しいMemoryRepositoryを作成します
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		reservations: make(map[string]*domain.Reservation),
	}
}

// Save は予約を保存します
func (r *MemoryRepository) Save(reservation *domain.Reservation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 重複チェック
	for _, existing := range r.reservations {
		if existing.Status == "booked" && reservation.Overlaps(existing) {
			return fmt.Errorf("時間が重複している予約が既に存在します")
		}
	}

	r.reservations[reservation.ID] = reservation
	return nil
}

// FindByID はIDで予約を検索します
func (r *MemoryRepository) FindByID(id string) (*domain.Reservation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	reservation, exists := r.reservations[id]
	if !exists {
		return nil, fmt.Errorf("予約が見つかりません: %s", id)
	}

	return reservation, nil
}

// FindAll はすべての予約を取得します
func (r *MemoryRepository) FindAll() ([]*domain.Reservation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*domain.Reservation, 0, len(r.reservations))
	for _, reservation := range r.reservations {
		result = append(result, reservation)
	}

	return result, nil
}

// Update は予約を更新します
func (r *MemoryRepository) Update(reservation *domain.Reservation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.reservations[reservation.ID]; !exists {
		return fmt.Errorf("予約が見つかりません: %s", reservation.ID)
	}

	r.reservations[reservation.ID] = reservation
	return nil
}

