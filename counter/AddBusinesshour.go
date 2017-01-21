package counter

import (
	//"errors"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// create a businesshour
// returns the inserted id and a slice of errors

func AddBusineshour(
	db *sql.DB,
	opening time.Time,
	closing time.Time,
) (
	int64,
	[]error,
) {

	var e error
	var err []error

	res, e := db.Exec(`
		INSERT INTO businesshour (
			opening,
			closing
		) VALUES (?, ?)`,
		opening,
		closing,
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
