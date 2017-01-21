package counter

import (
	"log"
	"time"
	//"errors"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// assign an existing businesshour to a weekday/ counter

func CreateBusinessHour(
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
		opening.Format("15:04:05"),
		closing.Format("15:04:05"),
	)
	if e != nil {
		log.Println(e)
		err = append(err, e)

		// query if entry already exist
		var id int
		e := db.QueryRow(`
			SELECT
				b.id
			FROM
				businesshour AS b
			WHERE
				b.opening = ?
				AND
				b.closing = ?`,
			opening,
			closing,
		).Scan(&id)
		if e != nil {
			log.Println(e)
		}
		return int64(id), nil
	}

	lastInsertID, e := res.LastInsertId()
	if e != nil {
		log.Println(e)
		err = append(err, e)
	}

	return lastInsertID, err

}
