package counter

import (
	"log"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// returns all business hours for a counter

func ListBusinessHoursForACounter(
	db *sql.DB,
	counterID int,
) []BusinessHour {

	var err []error
	var e error

	var businessHours []BusinessHour
	var b BusinessHour

	var opening string
	var closing string

	rows, e := db.Query(`
		SELECT
			b.id,
			cbw.weekday_id,
			b.opening,
			b.closing
		FROM
			 businesshour AS b
		JOIN
			counter_businesshour_weekday AS cbw ON cbw.businesshour_id = b.id
		WHERE
			cbw.counter_id = ?
		ORDER BY
			cbw.weekday_id ASC`,
		counterID,
	)
	defer rows.Close()

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

		rows.Scan(
			&b.ID,
			&b.Weekday,
			&opening,
			&closing,
		)
		b.Opening, e = time.Parse("15:04:05", opening)
		if e != nil {
			err = append(err, e)
		}
		b.Closing, e = time.Parse("15:04:05", closing)
		if e != nil {
			err = append(err, e)
		}
		businessHours = append(businessHours, b)
	}

	return businessHours

}
