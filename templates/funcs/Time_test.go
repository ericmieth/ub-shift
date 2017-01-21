package funcs

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	cases := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC), "20:34 Uhr"},
	}

	for _, c := range cases {
		got := Time(c.input)
		if got != c.want {
			t.Errorf("Time(%q) ==\n%q (is)\n%q (should)", c.input, got, c.want)
		}
	}
}
