package counter

import (
	"errors"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// insert a counter into the database
// returns the inserted id and a slice of errors

func AddCounter(
	db *sql.DB,
	name string,
	branch int,
) (
	int64,
	[]error,
) {

	var e error
	var err []error

	if len(name) == 0 {
		e = errors.New("Der Name der Theke darf nicht leer sein.")
		err = append(err, e)
	}

	// return if an error occurred
	if checkErrors(err) {
		return 0, err
	}

	res, e := db.Exec(`
		INSERT INTO counter (
			name,
			branch_id
		) VALUES (?, ?)`,
		name,
		branch,
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
