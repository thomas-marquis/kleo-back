package testing_utils

import (
	"time"
)

func NewDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0.0, 0.0, 0.0, 0.0, time.UTC)
}
