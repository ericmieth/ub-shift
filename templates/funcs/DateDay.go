package funcs

import (
	"time"
)

// returns the day of a date

func DateDay(t time.Time) int {
	return t.Day()
}
