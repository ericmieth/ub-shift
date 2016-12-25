package funcs

import (
	"time"
)

// returns the id of a weekday. Sunday = 0, ...

func DateWeekdayID(t time.Weekday) int {
	return int(t)
}
