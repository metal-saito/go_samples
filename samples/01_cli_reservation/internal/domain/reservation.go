package domain

import (
	"fmt"
	"time"
)

// Reservation は予約を表すドメインモデルです
type Reservation struct {
	ID           string
	UserName     string
	ResourceName string
	StartsAt     time.Time
	EndsAt       time.Time
	Status       string
}

// NewReservation は新しい予約を作成します
func NewReservation(userName, resourceName string, startsAt, endsAt time.Time) (*Reservation, error) {
	if userName == "" {
		return nil, fmt.Errorf("利用者名は必須です")
	}
	if resourceName == "" {
		return nil, fmt.Errorf("リソース名は必須です")
	}
	if endsAt.Before(startsAt) || endsAt.Equal(startsAt) {
		return nil, fmt.Errorf("終了時刻は開始時刻より後である必要があります")
	}

	return &Reservation{
		ID:           generateID(),
		UserName:     userName,
		ResourceName: resourceName,
		StartsAt:     startsAt,
		EndsAt:       endsAt,
		Status:       "booked",
	}, nil
}

// Overlaps は他の予約と時間が重複しているかチェックします
func (r *Reservation) Overlaps(other *Reservation) bool {
	if r.ResourceName != other.ResourceName {
		return false
	}
	return r.StartsAt.Before(other.EndsAt) && other.StartsAt.Before(r.EndsAt)
}

var idCounter = 1

func generateID() string {
	id := fmt.Sprintf("RES-%04d", idCounter)
	idCounter++
	return id
}

