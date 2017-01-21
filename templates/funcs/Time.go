package funcs

import (
	"time"
)

// FormatDate is formating a time to a string using an translation of
// month names
// if time is true, the date and time are returned
func Time(date time.Time) string {
	return date.Format("15:04") + " Uhr"
}
