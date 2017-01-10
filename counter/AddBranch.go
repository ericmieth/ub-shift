package counter

import (
	"errors"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// insert a branch into the database

func AddBranch(
	db *sql.DB,
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
		INSERT INTO branch (
			name,
			location
		) VALUES (?, ?)`,
		name,
		location,
	)

	err = append(err, e)
	return err
}