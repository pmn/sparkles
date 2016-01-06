package main

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

func initDBConn() *sql.DB {
	c := getConfig()

	connstr, err := pq.ParseURL(c.DBUri)

	if err != nil {
		log.Panic(err)
	}

	conn, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Panic(err)
	}

	return conn
}

func getSparklesForUser(user string) []Sparkle {
	db := initDBConn()
	defer db.Close()

	result, err := db.Query("SELECT Id, Sparkler, Sparklee, Reason, Room, Time FROM sparkles WHERE receiver = $1", user)
	defer result.Close()
	if err != nil {
		log.Panic(err.Error())
	}

	sparkles := []Sparkle{}

	for result.Next() {
		var s Sparkle
		result.Scan(&s.Id, &s.Sparklee, &s.Sparkler, &s.Reason, &s.Reason, &s.Room, &s.Time)

		sparkles = append(sparkles, s)
	}

	return sparkles
}

func getAllSparkles() []Sparkle {
	db := initDBConn()
	defer db.Close()

	result, err := db.Query("SELECT Id, Sparkler, Sparklee, Reason, Room, Time FROM sparkles")
	defer result.Close()
	if err != nil {
		log.Panic(err.Error())
	}

	sparkles := []Sparkle{}

	for result.Next() {
		var s Sparkle
		result.Scan(&s.Id, &s.Sparklee, &s.Sparkler, &s.Reason, &s.Reason, &s.Room, &s.Time)

		sparkles = append(sparkles, s)
	}

	return sparkles
}

func deleteSparkle(id int) {
	db := initDBConn()
	defer db.Close()

	_, err := db.Exec("DELETE FROM Sparkles WHERE Id = $1", id)
	if err != nil {
		log.Panic(err)
	}
}

func updateSparkle(id int, giver, reciever, reason, room, source string) {
	db := initDBConn()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Sparkles SET " +
		"Sparkler = $1, " +
		"Sparklee = $2, " +
		"Reason = $3, " +
		"Room = $4 " +
		"WHERE Id = $5")
	defer stmt.Close()
	if err != nil {
		log.Panic(err)
	}

	_, err = stmt.Exec(giver, reciever, reason, room, source, id)
	if err != nil {
		log.Panic(err)
	}
}
