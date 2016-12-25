package funcs

import (
	"log"
	"strconv"
	"time"
)

// return the following month

func DateNextMonth(year int, month time.Month) string {

	t, err := time.Parse(
		"2006-January",
		strconv.Itoa(year)+"-"+month.String(),
	)
	if err != nil {
		log.Println(err)
	}
	return t.AddDate(0, 1, 0).Format("2006-01")
}
