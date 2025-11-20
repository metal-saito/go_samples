package reservation

import (
	"fmt"
	"time"
)

// Reservation represents a reservation
type Reservation struct {
	ID           string
	UserName     string
	ResourceName string
	StartsAt     time.Time
	EndsAt       time.Time
	Status       string
}

// NewReservation creates a new reservation
func NewReservation(userName, resourceName string, startsAt, endsAt time.Time) (*Reservation, error) {
	if userName == "" {
		return nil, fmt.Errorf("user_name is required")
	}
	if resourceName == "" {
		return nil, fmt.Errorf("resource_name is required")
	}
	if endsAt.Before(startsAt) || endsAt.Equal(startsAt) {
		return nil, fmt.Errorf("ends_at must be after starts_at")
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

// Overlaps checks if this reservation overlaps with another
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

