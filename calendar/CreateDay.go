package calendar

import (
	"log"
	"time"

	"database/sql"
	_ "github.com/lib/pq"
)

// insert a row in table 'day'

func CreateDay(db *sql.DB, t time.Time) error {

	_, err := db.Exec(`
	INSERT INTO day (
		date,
		weekday_id
	)
	VALUES ($1, $2)`,
		t.Format("2006-01-02"),
		t.Weekday(),
	)

	if err != nil {
		log.Println(err)
	}

	return err
}
