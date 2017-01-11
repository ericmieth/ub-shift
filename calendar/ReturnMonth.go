package calendar

import (
	"log"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// returns all days belonging to a day
func ReturnMonth(
	db *sql.DB,
	year int,
	month time.Month,
) (
	[]Day,
	error,
) {

	var days []Day
	var d Day

	rows, err := db.Query(`
		SELECT
			d.id,
			d.date,
			d.weekday_id
		FROM
			day AS d
		WHERE
			YEAR(d.date) = ?
			AND
			MONTH(d.date) = ?
		ORDER BY
			d.date ASC`,
		year,
		int(month),
	)
	defer rows.Close()

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

		var dateString string
		rows.Scan(
			&d.ID,
			&dateString,
			&d.Weekday,
		)

		t, err := time.Parse("2006-01-02T15:04:05Z", dateString)
		if err != nil {
			log.Println(err)
		}
		d.Date = t

		days = append(days, d)
	}

	return days, err
}
