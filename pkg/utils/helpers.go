package utils

import (
	"fmt"
	"time"
)

// FormatDateTime formats a time.Time to ISO8601 string
func FormatDateTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ParseDateTime parses an ISO8601 string to time.Time
func ParseDateTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// ValidateTimeRange validates that end time is after start time
func ValidateTimeRange(startsAt, endsAt time.Time) error {
	if endsAt.Before(startsAt) || endsAt.Equal(startsAt) {
		return fmt.Errorf("ends_at must be after starts_at")
	}
	return nil
}

