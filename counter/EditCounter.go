package counter

import (
	"errors"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// insert a branch into the database

func EditCounter(
	db *sql.DB,
	counterID int,
	name string,
	branch int,
) []error {

	var e error
	var err []error

	if len(name) == 0 {
		e = errors.New("Der Name der Zweigstelle darf nicht leer sein.")
		err = append(err, e)
	}

	// return if an error occurred
	if checkErrors(err) {
		return err
	}

	_, e = db.Exec(`
		UPDATE
			counter
		SET
			name = ?,
			branch_id = ?
		WHERE
			id = ?`,
		name,
		branch,
		counterID,
	)

	err = append(err, e)
	return err
}
