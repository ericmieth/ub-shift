package counter

import (
	"errors"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// insert a branch into the database
// returns the inserted id and a slice of errors

func AddBranch(
	db *sql.DB,
	name string,
	location string,
) (
	int64,
	[]error,
) {

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
		return 0, err
	}

	res, e := db.Exec(`
		INSERT INTO branch (
			name,
			location
		) VALUES (?, ?)`,
		name,
		location,
	)
	if e != nil {
		err = append(err, e)
	}

	lastInsertID, e := res.LastInsertId()
	if e != nil {
		err = append(err, e)
		return 0, err
	}

	return lastInsertID, err

}
