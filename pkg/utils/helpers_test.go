package utils

import (
	"testing"
	"time"
)

func TestFormatDateTime(t *testing.T) {
	dt := time.Date(2025, 1, 2, 9, 0, 0, 0, time.UTC)
	formatted := FormatDateTime(dt)
	expected := "2025-01-02T09:00:00Z"
	
	if formatted != expected {
		t.Errorf("Expected %s, got %s", expected, formatted)
	}
}

func TestParseDateTime(t *testing.T) {
	dateStr := "2025-01-02T09:00:00Z"
	dt, err := ParseDateTime(dateStr)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	
	if dt.Year() != 2025 || dt.Month() != 1 || dt.Day() != 2 {
		t.Errorf("Parsed date is incorrect: %v", dt)
	}
}

func TestValidateTimeRange(t *testing.T) {
	startsAt := time.Date(2025, 1, 2, 9, 0, 0, 0, time.UTC)
	endsAt := time.Date(2025, 1, 2, 10, 0, 0, 0, time.UTC)
	
	err := ValidateTimeRange(startsAt, endsAt)
	if err != nil {
		t.Errorf("Expected no error for valid range, got %v", err)
	}
	
	err = ValidateTimeRange(endsAt, startsAt)
	if err == nil {
		t.Error("Expected error when ends_at is before starts_at")
	}
}

