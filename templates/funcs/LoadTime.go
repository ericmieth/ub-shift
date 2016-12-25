package funcs

import (
	"fmt"
	"time"
)

// returns the duration until now

func LoadTime(startTime time.Time) string {

	return fmt.Sprint(time.Since(startTime).Nanoseconds()/1e6) + "ms"
}
