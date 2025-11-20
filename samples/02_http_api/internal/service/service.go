package service

import (
	"fmt"
	"time"

	"reservation_api/internal/domain"
	"reservation_api/internal/repository"
)

// ReservationService は予約に関するビジネスロジックを提供します
type ReservationService struct {
	repo repository.Repository
}

// Repository はリポジトリインターフェースです
type Repository interface {
	Save(reservation *domain.Reservation) error
	FindByID(id string) (*domain.Reservation, error)
	FindAll() ([]*domain.Reservation, error)
	Update(reservation *domain.Reservation) error
}

// NewReservationService は新しいReservationServiceを作成します
func NewReservationService(repo repository.Repository) *ReservationService {
	return &ReservationService{repo: repo}
}

// CreateReservation は予約を作成します
func (s *ReservationService) CreateReservation(req *domain.CreateReservationRequest) (*domain.Reservation, error) {
	startsAt, err := time.Parse(time.RFC3339, req.StartsAt)
	if err != nil {
		return nil, fmt.Errorf("開始時刻の解析エラー: %w", err)
	}

	endsAt, err := time.Parse(time.RFC3339, req.EndsAt)
	if err != nil {
		return nil, fmt.Errorf("終了時刻の解析エラー: %w", err)
	}

	reservation, err := domain.NewReservation(req.UserName, req.ResourceName, startsAt, endsAt)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(reservation); err != nil {
		return nil, err
	}

	return reservation, nil
}

// ListReservations はすべての予約を取得します
func (s *ReservationService) ListReservations() ([]*domain.Reservation, error) {
	return s.repo.FindAll()
}

// CancelReservation は予約をキャンセルします
func (s *ReservationService) CancelReservation(id string) error {
	reservation, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	reservation.Status = "cancelled"
	return s.repo.Update(reservation)
}

