package calendar

import (
	"time"
)

// contains the date and weekday of a day
type Day struct {
	Date    time.Time    // date of this day
	Weekday time.Weekday // day of the week: Sun = 0, …
}
