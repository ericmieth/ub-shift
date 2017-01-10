package counter

import (
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// returns all branches

func ListCounters(
	db *sql.DB,
) (
	[]Counter,
	error,
) {

	var counters []Counter
	var c Counter

	rows, err := db.Query(`
		SELECT
			c.id,
			c.name,
			b.name,
			b.location
		FROM
			counter AS c
		JOIN
			branch AS b ON c.branch_id = b.id
		ORDER BY
			b.id ASC,
			c.name ASC`,
	)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return counters, err
	}

	for rows.Next() {

		rows.Scan(
			&c.ID,
			&c.Name,
			&c.Branch.Name,
			&c.Branch.Location,
		)

		counters = append(counters, c)
	}

	return counters, err
}
