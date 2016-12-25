package funcs

import (
	"github.com/ericmieth/ub-shift/calendar"
)

// fills a slice of days with some days at begin and end so there is
// an appropriate slice for rendering in template

func DateFillMonth(days []calendar.Day) []calendar.Day {

	firstDayOfMonth := days[0]
	firstWeekdayID := int(firstDayOfMonth.Weekday)

	lastDayOfMonth := days[len(days)-1]
	lastWeekdayID := int(lastDayOfMonth.Weekday)

	var offsetBegin []calendar.Day
	var offsetEnd []calendar.Day

	// handle offset at begin of a month

	// handle the case, if a week starts at sunday
	if firstWeekdayID == 0 {
		firstWeekdayID = 7
	}

	for i := 1; i < firstWeekdayID; i++ {
		date := firstDayOfMonth.Date.AddDate(
			0, 0, (firstWeekdayID-i)*(-1),
		)

		d := calendar.Day{
			Date:    date,
			Weekday: date.Weekday(),
		}
		offsetBegin = append(offsetBegin, d)
	}

	// handle offset at end of a month

	offsetEndCounter := 1
	for i := lastWeekdayID; ; i++ {
		if i%7 == 0 {
			break
		}
		date := lastDayOfMonth.Date.AddDate(
			0, 0, offsetEndCounter,
		)

		d := calendar.Day{
			Date:    date,
			Weekday: date.Weekday(),
		}

		offsetEndCounter++
		offsetEnd = append(offsetEnd, d)
	}

	days = append(offsetBegin, days...)
	days = append(days, offsetEnd...)

	return days
}
