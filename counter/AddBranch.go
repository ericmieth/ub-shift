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
) error {

	if len(name) == 0 {
		return errors.New("Der Name der Zweigstelle darf nicht leer sein.")
	}

	if len(location) == 0 {
		return errors.New("Die Adresse der Zweigstelle darf nicht leer sein.")
	}
	_, err := db.Exec(`
		INSERT INTO branch (
			name,
			location
		) VALUES (?, ?)`,
		name,
		location,
	)

	if err != nil {
		return err
	}

	return err
}
