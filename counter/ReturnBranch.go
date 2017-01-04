package counter

import (
	//"errors"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// insert a branch into the database

func ReturnBranch(
	db *sql.DB,
	id int,
) (
	Branch,
	error,
) {

	var b Branch

	err := db.QueryRow(`
		SELECT
			b.id,
			b.name,
			b.location
		FROM
			branch AS b
		WHERE
			b.id = ?  `,
		id,
	).Scan(
		&b.ID,
		&b.Name,
		&b.Location,
	)

	return b, err
}
