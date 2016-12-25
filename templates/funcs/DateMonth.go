package funcs

import (
	"time"
)

// returns the month of a date

func DateMonth(t time.Time) time.Month {
	return t.Month()
}
