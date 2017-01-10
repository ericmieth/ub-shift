package counter

import (
	"errors"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// insert a counter into the database

func AddCounter(
	db *sql.DB,
	name string,
	branch int,
) []error {

	var e error
	var err []error

	if len(name) == 0 {
		e = errors.New("Der Name der Theke darf nicht leer sein.")
		err = append(err, e)
	}

	// return if an error occurred
	if checkErrors(err) {
		return err
	}

	_, e = db.Exec(`
		INSERT INTO counter (
			name,
			branch_id
		) VALUES (?, ?)`,
		name,
		branch,
	)

	err = append(err, e)
	return err
}
