package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"sort"
	"time"
)

// HA HA, joke's on you! ENTIRE DB IS FILE!
const filename = "sparkledb"

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
		log.Print("There was an error loading the sparkle database. Using a blank one.")
	}

	return sparkleDB
}

func (sparkledb *SparkleDatabase) AddSparkle(sparkle Sparkle) Leader {
	// Add a sparkle to the database
	sparkle.Time = time.Now()
	sparkledb.Sparkles = append(sparkledb.Sparkles, sparkle)
	giver := Leader{Name: sparkle.Sparkler, Score: 1}
	receiver := Leader{Name: sparkle.Sparklee, Score: 1}

	// Add the receiver's data
	receiver_found := false
	for k, v := range sparkledb.MostReceived {
		if v.Name == sparkle.Sparklee {
			receiver_found = true
			sparkledb.MostReceived[k].Score++
			receiver.Score = sparkledb.MostReceived[k].Score
		}
	}

	// Add the receiver if not already there
	if !receiver_found {
		sparkledb.MostReceived = append(sparkledb.MostReceived, receiver)
	}

	// Add the giver's data
	giver_found := false
	for k, v := range sparkledb.MostGiven {
		if v.Name == sparkle.Sparkler {
			giver_found = true
			sparkledb.MostGiven[k].Score++
		}
	}

	// Add the giver if not already there
	if !giver_found {
		sparkledb.MostGiven = append(sparkledb.MostGiven, giver)
	}

	// After the sparkle has been added, save the data file
	sparkledb.Save()

	// Return the receiver record so that Hubot can report the users total sparkles
	return receiver
}

type ByScore []Leader

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

func (sparkledb *SparkleDatabase) TopReceivers(n int) []Leader {
	// Get the top N leaders and return them
	sort.Sort(sort.Reverse(ByScore(sparkledb.MostReceived)))
	return sparkledb.MostReceived
}

func (sparkledb *SparkleDatabase) TopGivers(n int) []Leader {
	// Get the top N leaders and return them
	sort.Sort(sort.Reverse(ByScore(sparkledb.MostGiven)))
	return sparkledb.MostGiven
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
