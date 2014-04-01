package main

// HA HA, joke's on you! ENTIRE DB IS FILE!
const filename = "sparkledb.gob"

type SparkleDatabase struct {
	Sparkles []Sparkle
}

func (db *SparkleDatabase) Save() {
	// Persist the database to file

}

func (db *SparkleDatabase) Load() {
	// Load the database from a file
}

func (db *SparkleDatabase) AddSparkle(sparkle Sparkle) {
	// Add a sparkle to the database
}

func (db *SparkleDatabase) TopSparkles(n int) []Sparkle {
	// Return top n sparkles

	return []Sparkle{}
}

func (db *SparkleDatabase) SparklesForUser(user string) []Sparkle {
	// Return all the sparkles for <user>
	return []Sparkle{}
}
