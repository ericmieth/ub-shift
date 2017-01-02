package employee

import (
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// returns all days belonging to a day
func ListEmployees(
	db *sql.DB,
) (
	[]Employee,
	error,
) {

	var emplyees []Employee
	var e Employee

	rows, err := db.Query(`
		SELECT
			e.id,
			e.firstname,
			e.lastname,
			e.manager,
			e.mailaddress,
			e.active,
			e.workinghours
		FROM
			employee AS e
		ORDER BY
			e.lastname ASC`,
	)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return emplyees, err
	}

	for rows.Next() {

		rows.Scan(
			&e.ID,
			&e.FirstName,
			&e.LastName,
			&e.Manager,
			&e.MailAddress,
			&e.Active,
			&e.WorkingHours,
		)

		log.Println(e)

		emplyees = append(emplyees, e)
	}

	return emplyees, err
}
