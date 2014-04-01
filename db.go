package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

// HA HA, joke's on you! ENTIRE DB IS FILE!
const filename = "sparkledb.gob"

type SparkleDatabase struct {
	Sparkles     []Sparkle
	MostReceived []Leader
	MostGiven    []Leader
}

func (sparkledb *SparkleDatabase) Save() {
	// Persist the database to file
	var data bytes.Buffer
	contents := gob.NewEncoder(&data)
	err := contents.Encode(sparkledb)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename, data.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func LoadDB() SparkleDatabase {
	// Load the database from a file
	n, err := ioutil.ReadFile(filename)
	// Create a bytes.Buffer
	p := bytes.NewBuffer(n)
	dec := gob.NewDecoder(p)

	var sparkleDB SparkleDatabase
	err = dec.Decode(&sparkleDB)

	if err != nil {
		panic(err)
	}

	return sparkleDB
}

func (sparkledb *SparkleDatabase) AddSparkle(sparkle Sparkle) {
	// Add a sparkle to the database
	sparkledb.Sparkles = append(sparkledb.Sparkles, sparkle)
	giver := Leader{Name: sparkle.Sparkler, Score: 1}
	receiver := Leader{Name: sparkle.Sparklee, Score: 1}

	receiver_found := false
	for k, v := range sparkledb.MostReceived {
		if v.Name == sparkle.Sparklee {
			receiver_found = true
			sparkledb.MostReceived[k].Score++
		}
	}

	if !receiver_found {
		sparkledb.MostReceived = append(sparkledb.MostReceived, receiver)
	}

	giver_found := false
	for k, v := range sparkledb.MostGiven {
		if v.Name == sparkle.Sparkler {
			giver_found = true
			sparkledb.MostGiven[k].Score++
		}
	}

	if !giver_found {
		sparkledb.MostGiven = append(sparkledb.MostGiven, giver)
	}

	// After the sparkle has been added, save the data file
	sparkledb.Save()
}

func (sparkledb *SparkleDatabase) TopUsers(n int) []Leader {
	// Get the top N leaders and return them
	leaders := make([]Leader, n)

	return leaders
}

func (db *SparkleDatabase) SparklesForUser(user string) []Sparkle {
	// Return all the sparkles for <user>
	var list []Sparkle
	for _, v := range db.Sparkles {
		if v.Sparklee == user {
			list = append(list, v)
		}
	}

	return list
}
