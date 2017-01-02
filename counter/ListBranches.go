package counter

import (
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// returns all branches

func ListBranches(
	db *sql.DB,
) (
	[]Branch,
	error,
) {

	var branches []Branch
	var b Branch

	rows, err := db.Query(`
		SELECT
			b.id,
			b.name,
			b.location
		FROM
			branch AS b
		ORDER BY
			b.id ASC`,
	)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return branches, err
	}

	for rows.Next() {

		rows.Scan(
			&b.ID,
			&b.Name,
			&b.Location,
		)

		branches = append(branches, b)
	}

	return branches, err
}
