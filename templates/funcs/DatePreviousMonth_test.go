package funcs

import (
	"testing"
	"time"
)

func TestDatePreviousMonth(t *testing.T) {
	cases := []struct {
		inputYear  int
		inputMonth time.Month
		want       string
	}{
		{2009, 11, "2009-10"},
		{2009, 01, "2008-12"},
	}

	for _, c := range cases {
		got := DatePreviousMonth(c.inputYear, c.inputMonth)
		if got != c.want {
			t.Errorf(
				"DatePreviousMonth(%q,%q) ==\n%q (is)\n%q (should)",
				c.inputYear,
				c.inputMonth,
				got,
				c.want,
			)
		}
	}

}
