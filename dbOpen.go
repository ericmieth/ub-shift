package main

import (
	"log"
	"strings"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func dbOpen() *sql.DB {
	c := readConfig()

	config := []string{
		c["DBUser"], ":",
		c["DBPass"], "@tcp(",
		c["DBHost"], ":",
		c["DBPort"], ")/",
		c["DBName"], "?parseTime=true",
	}

	db, err := sql.Open("mysql", strings.Join(config, ""))
	if err != nil {
		log.Print(err)
	}
	return db
}
