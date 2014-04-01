package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

// HA HA, joke's on you! ENTIRE DB IS FILE!
const filename = "sparkledb.gob"

type SparkleDatabase struct {
	Sparkles []Sparkle
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
	// After the sparkle has been added, save the data file
	sparkledb.Save()
}

func (sparkledb *SparkleDatabase) TopUsers(n int) []Sparkle {
	// Return users with the most sparkles

	return []Sparkle{}
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
