package main

import (
	"database/sql"
	"log"
	"strings"
)

func dbOpen() *sql.DB {
	c := readConfig()

	config := []string{
		"host=", c["DBHost"],
		"port=", c["DBPort"],
		"dbname=", c["DBName"],
		"user=", c["DBUser"],
		"password=", c["DBPass"],
		"sslmode=disable",
	}

	db, err := sql.Open("postgres", strings.Join(config, " "))
	if err != nil {
		log.Print(err)
	}
	return db
}
