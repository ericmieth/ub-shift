package calendar

import (
	"log"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// insert a row in table 'day'

func CreateDay(db *sql.DB, t time.Time) error {

	_, err := db.Exec(`
	INSERT INTO day (
		date,
		weekday_id
	)
	VALUES (?, ?)`,
		t.Format("2006-01-02"),
		t.Weekday(),
	)

	if err != nil {
		log.Println(err)
	}

	return err
}
