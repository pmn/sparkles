package main

import "os"

type config struct {
	DBUri string
}

func getConfig() config {
	var c config

	dburi := os.Getenv("KARMA_DB_URL")

	if dburi == "" {
		c.DBUri = "postgres://karma:karma@localhost/karma"
	} else {
		c.DBUri = dburi
	}

	return c
}
