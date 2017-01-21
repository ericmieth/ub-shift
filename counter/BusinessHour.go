package counter

import (
	"time"
)

// contains information about a business hour for a counter

type BusinessHour struct {
	ID      int          // ID
	Weekday time.Weekday // day of the week: Sun = 0, …
	Opening time.Time    // begin of a business hour
	Closing time.Time    // end of a business hour
}
