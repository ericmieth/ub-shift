package counter

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ReturnBusinessHoursForAllDays(
	db *sql.DB,
	counterID int,
) (
	[]BusinessHoursForADay,
	[]error,
) {

	var e error
	var err []error

	var days []BusinessHoursForADay

	// iterate over all days, starting with monday
	//for i := 0; i < 7; i++ {
	for weekdayID := 0; ; {

		weekdayID += 1
		weekdayID %= 7

		var d BusinessHoursForADay

		// get name of the day

		e = db.QueryRow(`
			SELECT
				w.name
			FROM
				weekday AS w
			WHERE
				w.id = ?  `,
			weekdayID,
		).Scan(
			&d.Name,
		)

		if e != nil {
			err = append(err, e)
		}

		d.ID = weekdayID

		// get assigned businesshours

		assignedBusinessHours, err2 := ReturnAssignedBusinessHoursForADay(
			db,
			counterID,
			weekdayID,
		)
		d.AssignedBusinessHours = assignedBusinessHours
		if checkErrors(err2) {
			err = append(err, err2...)
		}

		// get existing businesshours without the assigned ones
		// (remaining ones)

		remainingBusinessHours, err2 := ReturnRemainingBusinessHoursForADay(
			db,
			counterID,
			weekdayID,
		)
		d.RemainingBusinessHours = remainingBusinessHours
		if checkErrors(err2) {
			err = append(err, err2...)
		}

		days = append(days, d)

		if weekdayID == 0 {
			break
		}

	}

	return days, err
}
