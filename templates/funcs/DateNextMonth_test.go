package funcs

import (
	"testing"
	"time"
)

func TestDateNextMonth(t *testing.T) {
	cases := []struct {
		inputYear  int
		inputMonth time.Month
		want       string
	}{
		{2009, 11, "2009-12"},
		{2009, 12, "2010-01"},
	}

	for _, c := range cases {
		got := DateNextMonth(c.inputYear, c.inputMonth)
		if got != c.want {
			t.Errorf(
				"DateNextMonth(%q,%q) ==\n%q (is)\n%q (should)",
				c.inputYear,
				c.inputMonth,
				got,
				c.want,
			)
		}
	}

}
