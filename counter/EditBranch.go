package counter

import (
	"errors"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// insert a branch into the database

func EditBranch(
	db *sql.DB,
	id int,
	name, location string,
) []error {

	var e error
	var err []error

	if len(name) == 0 {
		e = errors.New("Der Name der Zweigstelle darf nicht leer sein.")
		err = append(err, e)
	}

	if len(location) == 0 {
		e = errors.New("Die Adresse der Zweigstelle darf nicht leer sein.")
		err = append(err, e)
	}

	// return if an error occurred
	if checkErrors(err) {
		return err
	}

	_, e = db.Exec(`
		UPDATE branch
		SET name = ?,
			location = ?
		WHERE id = ?`,
		name,
		location,
		id,
	)

	err = append(err, e)
	return err
}
