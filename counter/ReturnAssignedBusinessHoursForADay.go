package counter

import (
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// returns the assigned business hours for a day on a counter

func ReturnAssignedBusinessHoursForADay(
	db *sql.DB,
	counterID int,
	weekdayID int,
) (
	[]BusinessHour,
	[]error,
) {

	var e error
	var err []error

	var assignedBusinessHours []BusinessHour
	var b BusinessHour

	var opening string
	var closing string

	rows, e := db.Query(`
			SELECT
				b.id,
				b.opening,
				b.closing
			FROM
				counter_businesshour_weekday AS cbw
			JOIN businesshour AS b ON cbw.businesshour_id = b.id
			WHERE
				cbw.counter_id = ?
				AND
				cbw.weekday_id = ?`,
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

		assignedBusinessHours = append(assignedBusinessHours, b)
	}

	return assignedBusinessHours, err
}
