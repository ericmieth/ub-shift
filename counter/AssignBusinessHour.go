package counter

import (
	"errors"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// assign an existing businesshour to a weekday/ counter

func AssignBusinessHour(
	db *sql.DB,
	counterID int,
	weekdayID int,
	businesshourID int,
) []error {

	var e error
	var err []error

	// check against valid weekday id

	if weekdayID < 0 || weekdayID > 6 {
		e = errors.New("ungültiger Wochentag.")
		err = append(err, e)
	}

	// return if an error occurred
	if checkErrors(err) {
		return err
	}

	_, e = db.Exec(`
		INSERT INTO counter_businesshour_weekday (
			counter_id,
			businesshour_id,
			weekday_id
		) VALUES (?, ?, ?)`,
		counterID,
		businesshourID,
		weekdayID,
	)
	if e != nil {
		err = append(err, e)
	}

	return err

}
