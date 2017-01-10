package counter

import (
	//"errors"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// insert a branch into the database

func ReturnCounter(
	db *sql.DB,
	id int,
) (
	Counter,
	error,
) {

	var c Counter

	err := db.QueryRow(`
		SELECT
			c.id,
			c.name,
			b.id,
			b.name,
			b.location
		FROM
			counter AS c
		JOIN 
			branch AS b ON b.id = c.branch_id
		WHERE
			c.id = ?  `,
		id,
	).Scan(
		&c.ID,
		&c.Name,
		&c.Branch.ID,
		&c.Branch.Name,
		&c.Branch.Location,
	)

	return c, err
}
