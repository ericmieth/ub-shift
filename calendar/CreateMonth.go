package calendar

import (
	"database/sql"
	"log"
	"strconv"
	"time"
)

// create a databae entry for each day of a month

func CreateMonth(db *sql.DB, year int, month time.Month) error {

	date := strconv.Itoa(year) + "-" + month.String()
	t, err := time.Parse("2006-January", date)

	if err != nil {
		log.Println(err)
	}

	for {
		if t.Month() != month {
			break
		}
		CreateDay(db, t)
		t = t.AddDate(0, 0, 1)
	}

	return err
}
