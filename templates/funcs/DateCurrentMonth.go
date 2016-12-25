package funcs

import (
	"time"
)

// returns the current month

func DateCurrentMonth() string {

	return time.Now().Format("2006-01")
}
