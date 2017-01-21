package counter

import (
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// get existing businesshours without the assigned ones
// (remaining ones)

func ReturnRemainingBusinessHoursForADay(
	db *sql.DB,
	counterID int,
	weekdayID int,
) (
	[]BusinessHour,
	[]error,
) {
	var e error
	var err []error

	var remainingBusinessHours []BusinessHour
	var b BusinessHour

	var opening string
	var closing string

	rows, e := db.Query(`
		SELECT
			DISTINCT b.id,
			b.opening,
			b.closing
		FROM
			businesshour AS b
		LEFT JOIN
			counter_businesshour_weekday AS cbw
			ON cbw.businesshour_id = b.id
		WHERE
			b.id
			NOT IN (
				SELECT
					cbw.businesshour_id
				FROM
					counter_businesshour_weekday AS cbw
				WHERE
					cbw.counter_id = ?
					AND
					cbw.weekday_id = ?
			)`,
		counterID,
		weekdayID,
	)

	defer rows.Close()

	if e != nil {
		err = append(err, e)
	}

	for rows.Next() {

		rows.Scan(
			&b.ID,
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

		remainingBusinessHours = append(remainingBusinessHours, b)
	}

	return remainingBusinessHours, err

}
