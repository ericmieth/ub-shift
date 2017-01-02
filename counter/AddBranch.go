package counter

import (
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// insert a branch into the database

func AddBranch(
	db *sql.DB,
	name, location string,
) error {

	_, err := db.Exec(`
		INSERT INTO branch (
			name,
			location
		) VALUES (?, ?)`,
		name,
		location,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
