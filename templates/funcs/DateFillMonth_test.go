package funcs

import (
	"github.com/ericmieth/ub-shift/calendar"

	"log"
	"testing"
	"time"
)

// compares two slices containing calendar.Day
func compareDaySlices(d1 []calendar.Day, d2 []calendar.Day) bool {
	if len(d1) != len(d2) {
		log.Println(len(d1))
		return false
	}
	for i, _ := range d1 {
		if d1[i].Date != d2[i].Date || d1[i].Weekday != d2[i].Weekday {
			log.Println(d1[i])
			log.Println(d2[i])
			return false
		}
	}
	return true
}

func TestDateFillMonth(t *testing.T) {

	cases := []struct {
		input []calendar.Day
		want  []calendar.Day
	}{
		// case 1
		{
			// input
			[]calendar.Day{
				calendar.Day{
					Date:    time.Date(2009, 11, 1, 00, 00, 00, 000000000, time.UTC),
					Weekday: 0,
				},
				calendar.Day{
					Date:    time.Date(2009, 11, 2, 00, 00, 00, 000000000, time.UTC),
					Weekday: 1,
				},
			},
			// want
			[]calendar.Day{
				calendar.Day{
					Date:    time.Date(2009, 10, 26, 00, 00, 00, 000000000, time.UTC),
					Weekday: 1,
				},
				calendar.Day{
					Date:    time.Date(2009, 10, 27, 00, 00, 00, 000000000, time.UTC),
					Weekday: 2,
				},
				calendar.Day{
					Date:    time.Date(2009, 10, 28, 00, 00, 00, 000000000, time.UTC),
					Weekday: 3,
				},
				calendar.Day{
					Date:    time.Date(2009, 10, 29, 00, 00, 00, 000000000, time.UTC),
					Weekday: 4,
				},
				calendar.Day{
					Date:    time.Date(2009, 10, 30, 00, 00, 00, 000000000, time.UTC),
					Weekday: 5,
				},
				calendar.Day{
					Date:    time.Date(2009, 10, 31, 00, 00, 00, 000000000, time.UTC),
					Weekday: 6,
				},
				calendar.Day{
					Date:    time.Date(2009, 11, 1, 00, 00, 00, 000000000, time.UTC),
					Weekday: 0,
				},
				calendar.Day{
					Date:    time.Date(2009, 11, 2, 00, 00, 00, 000000000, time.UTC),
					Weekday: 1,
				},
				calendar.Day{
					Date:    time.Date(2009, 11, 3, 00, 00, 00, 000000000, time.UTC),
					Weekday: 2,
				},
				calendar.Day{
					Date:    time.Date(2009, 11, 4, 00, 00, 00, 000000000, time.UTC),
					Weekday: 3,
				},
				calendar.Day{
					Date:    time.Date(2009, 11, 5, 00, 00, 00, 000000000, time.UTC),
					Weekday: 4,
				},
				calendar.Day{
					Date:    time.Date(2009, 11, 6, 00, 00, 00, 000000000, time.UTC),
					Weekday: 5,
				},
				calendar.Day{
					Date:    time.Date(2009, 11, 7, 00, 00, 00, 000000000, time.UTC),
					Weekday: 6,
				},
				calendar.Day{
					Date:    time.Date(2009, 11, 8, 00, 00, 00, 000000000, time.UTC),
					Weekday: 0,
				},
			},
		},
		// case 2
		{
			// input
			[]calendar.Day{
				calendar.Day{
					Date:    time.Date(2009, 8, 1, 00, 00, 00, 000000000, time.UTC),
					Weekday: 6,
				},
				calendar.Day{
					Date:    time.Date(2009, 8, 2, 00, 00, 00, 000000000, time.UTC),
					Weekday: 0,
				},
				calendar.Day{
					Date:    time.Date(2009, 8, 3, 00, 00, 00, 000000000, time.UTC),
					Weekday: 1,
				},
				calendar.Day{
					Date:    time.Date(2009, 8, 4, 00, 00, 00, 000000000, time.UTC),
					Weekday: 2,
				},
			},
			// want
			[]calendar.Day{
				calendar.Day{
					Date:    time.Date(2009, 7, 27, 00, 00, 00, 000000000, time.UTC),
					Weekday: 1,
				},
				calendar.Day{
					Date:    time.Date(2009, 7, 28, 00, 00, 00, 000000000, time.UTC),
					Weekday: 2,
				},
				calendar.Day{
					Date:    time.Date(2009, 7, 29, 00, 00, 00, 000000000, time.UTC),
					Weekday: 3,
				},
				calendar.Day{
					Date:    time.Date(2009, 7, 30, 00, 00, 00, 000000000, time.UTC),
					Weekday: 4,
				},
				calendar.Day{
					Date:    time.Date(2009, 7, 31, 00, 00, 00, 000000000, time.UTC),
					Weekday: 5,
				},
				calendar.Day{
					Date:    time.Date(2009, 8, 1, 00, 00, 00, 000000000, time.UTC),
					Weekday: 6,
				},
				calendar.Day{
					Date:    time.Date(2009, 8, 2, 00, 00, 00, 000000000, time.UTC),
					Weekday: 0,
				},
				calendar.Day{
					Date:    time.Date(2009, 8, 3, 00, 00, 00, 000000000, time.UTC),
					Weekday: 1,
				},
				calendar.Day{
					Date:    time.Date(2009, 8, 4, 00, 00, 00, 000000000, time.UTC),
					Weekday: 2,
				},
				calendar.Day{
					Date:    time.Date(2009, 8, 5, 00, 00, 00, 000000000, time.UTC),
					Weekday: 3,
				},
				calendar.Day{
					Date:    time.Date(2009, 8, 6, 00, 00, 00, 000000000, time.UTC),
					Weekday: 4,
				},
				calendar.Day{
					Date:    time.Date(2009, 8, 7, 00, 00, 00, 000000000, time.UTC),
					Weekday: 5,
				},
				calendar.Day{
					Date:    time.Date(2009, 8, 8, 00, 00, 00, 000000000, time.UTC),
					Weekday: 6,
				},
				calendar.Day{
					Date:    time.Date(2009, 8, 9, 00, 00, 00, 000000000, time.UTC),
					Weekday: 0,
				},
			},
		},
	}

	for _, c := range cases {
		got := DateFillMonth(c.input)
		if !compareDaySlices(got, c.want) {
			t.Errorf(
				"DateFillMonth(%q) ==\n%q (is)\n%q (should)",
				c.input,
				got,
				c.want,
			)
		}
	}

}
