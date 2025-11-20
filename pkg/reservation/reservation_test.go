package reservation

import (
	"testing"
	"time"
)

func TestNewReservation(t *testing.T) {
	startsAt := time.Now().Add(24 * time.Hour)
	endsAt := startsAt.Add(time.Hour)

	reservation, err := NewReservation("Alice", "Room-A", startsAt, endsAt)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if reservation.UserName != "Alice" {
		t.Errorf("Expected UserName to be 'Alice', got %s", reservation.UserName)
	}

	if reservation.ResourceName != "Room-A" {
		t.Errorf("Expected ResourceName to be 'Room-A', got %s", reservation.ResourceName)
	}
}

func TestNewReservation_Validation(t *testing.T) {
	startsAt := time.Now().Add(24 * time.Hour)
	endsAt := startsAt.Add(time.Hour)

	_, err := NewReservation("", "Room-A", startsAt, endsAt)
	if err == nil {
		t.Error("Expected error for empty user_name")
	}

	_, err = NewReservation("Alice", "", startsAt, endsAt)
	if err == nil {
		t.Error("Expected error for empty resource_name")
	}

	_, err = NewReservation("Alice", "Room-A", endsAt, startsAt)
	if err == nil {
		t.Error("Expected error when ends_at is before starts_at")
	}
}

func TestReservation_Overlaps(t *testing.T) {
	baseTime := time.Now().Add(24 * time.Hour)
	r1, _ := NewReservation("Alice", "Room-A", baseTime, baseTime.Add(time.Hour))
	r2, _ := NewReservation("Bob", "Room-A", baseTime.Add(30*time.Minute), baseTime.Add(90*time.Minute))

	if !r1.Overlaps(r2) {
		t.Error("Expected reservations to overlap")
	}

	r3, _ := NewReservation("Charlie", "Room-B", baseTime, baseTime.Add(time.Hour))
	if r1.Overlaps(r3) {
		t.Error("Expected reservations not to overlap (different resources)")
	}
}

