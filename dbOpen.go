package main

import (
	"log"
	"strings"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func dbOpen(c Config) *sql.DB {

	config := []string{
		c.Database.User, ":",
		c.Database.Pass, "@tcp(",
		c.Database.Host, ":",
		c.Database.Port, ")/",
		c.Database.Name, "?parseTime=true",
	}

	db, err := sql.Open("mysql", strings.Join(config, ""))
	if err != nil {
		log.Print(err)
	}
	return db
}
