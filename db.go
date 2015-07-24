package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"sort"
	"strings"
	"time"

	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

// HA HA, joke's on you! ENTIRE DB IS FILE!
const filename = "sparkledb"
const bucketName = "mister-sparkleo"

// SparkleDatabase holds all the sparkle data
type SparkleDatabase struct {
	Sparkles []Sparkle
}

// Save the database
func (s *SparkleDatabase) Save() {
	// Persist the database to file
	var data bytes.Buffer
	contents := gob.NewEncoder(&data)
	err := contents.Encode(s)
	if err != nil {
		panic(err)
	}

	// The AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables are used.
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err.Error())
	}

	// Open Bucket
	s := s3.New(auth, aws.USEast)

	// Load the database from an S3 bucket
	bucket := s.Bucket(bucketName)

	err = bucket.Put(filename, data.Bytes(), "text/plain", s3.BucketOwnerFull)
	if err != nil {
		panic(err.Error())
	}
}

// LoadDB loads the SparkleDatabase from S3
func LoadDB() SparkleDatabase {
	// The AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables are used.
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err.Error())
	}

	// Open Bucket
	s := s3.New(auth, aws.USEast)

	// Load the database from an S3 bucket
	bucket := s.Bucket(bucketName)

	// Create a bytes.Buffer
	n, err := bucket.Get(filename)
	if err != nil {
		panic(err)
	}

	p := bytes.NewBuffer(n)
	dec := gob.NewDecoder(p)

	var sparkleDB SparkleDatabase
	err = dec.Decode(&sparkleDB)

	if err != nil {
		log.Print("There was an error loading the sparkle database. Using a blank one.")
	}

	return sparkleDB
}

// AddSparkle adds a sparkle to the database and returns a Leader record
func (s *SparkleDatabase) AddSparkle(sparkle Sparkle) Leader {
	// Add a sparkle to the database
	sparkle.Time = time.Now()
	s.Sparkles = append(sparkledb.Sparkles, sparkle)

	// After the sparkle has been added, save the data file
	s.Save()

	// Return the receiver record so that Hubot can report the users total sparkles
	var t time.Time
	receivers := sparkledb.Receivers(t)
	var recipient Leader
	for _, v := range receivers {
		if v.Name == sparkle.Sparklee {
			recipient = v
		}
	}

	return recipient
}

// Givers returns the top Leaders
func (s *SparkleDatabase) Givers(earliestDate time.Time) []Leader {
	var g = make(map[string]int)
	for _, v := range s.Sparkles {
		if v.Time.After(earliestDate) {
			g[v.Sparkler]++
		}
	}

	var leaders []Leader
	for k, v := range g {
		leader := Leader{Name: k, Score: v}
		leaders = append(leaders, leader)
	}

	return leaders
}

// Receivers returns the top Receivers
func (s *SparkleDatabase) Receivers(earliestDate time.Time) []Leader {
	var g = make(map[string]int)
	for _, v := range s.Sparkles {
		if v.Time.After(earliestDate) {
			g[v.Sparklee]++
		}
	}

	var leaders []Leader
	for k, v := range g {
		leader := Leader{Name: k, Score: v}
		leaders = append(leaders, leader)
	}

	return leaders
}

// ByScore is used for building the leaderboard
type ByScore []Leader

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

// TopGiven returns the top Givers
func (s *SparkleDatabase) TopGiven(since time.Time) []Leader {
	leaders := s.Givers(since)
	sort.Sort(sort.Reverse(ByScore(leaders)))
	return leaders
}

// TopReceived returns the top Receivers
func (s *SparkleDatabase) TopReceived(since time.Time) []Leader {
	leaders := s.Receivers(since)
	sort.Sort(sort.Reverse(ByScore(leaders)))
	return leaders
}

// SparklesForUser returns sparkles for user <user>
func (s *SparkleDatabase) SparklesForUser(user string) []Sparkle {
	// Return all the sparkles for <user>
	var list []Sparkle
	for _, v := range s.Sparkles {
		if strings.ToLower(v.Sparklee) == strings.ToLower(user) {
			list = append(list, v)
		}
	}

	return list
}
