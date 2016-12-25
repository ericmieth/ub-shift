package main

import (
	"encoding/json"
	"log"
	"os"
)

func readConfig() map[string]string {

	type config struct {
		DBHost string
		DBPort string
		DBName string
		DBUser string
		DBPass string
	}

	file, err := os.Open("config.json")
	decoder := json.NewDecoder(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	configuration := config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	c := make(map[string]string)
	c["DBHost"] = configuration.DBHost
	c["DBPort"] = configuration.DBPort
	c["DBName"] = configuration.DBName
	c["DBUser"] = configuration.DBUser
	c["DBPass"] = configuration.DBPass

	return c

}
