package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Sparkle datatype
type Sparkle struct {
	Id       int       `json:"id"`
	Sparkler string    `json:"sparkler"`
	Sparklee string    `json:"sparklee"`
	Reason   string    `json:"reason,omitempty"`
	Room     string    `json:"room,omitempty"`
	Time     time.Time `json:"time,omitempty"`
}

// Leader datatype
type Leader struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// Return the data in JSON format. This is the default return method.
func returnJSON(obj interface{}, w http.ResponseWriter, h *http.Request) {
	// Don't cache json returns. This is to work around ie's weird caching behavior
	w.Header().Set("Cache-Control", "no-cache")
	// Set the content type to json
	w.Header().Set("Content-Type", "application/json")

	// CORS to allow all callers
	w.Header().Set("Access-Control-Allow-Origin", "*")

	j, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprint(w, string(j))
}

// Placeholder in case anyone calls the root of the service.
// Perhaps change this to 404.
func defaultHandler(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Default sparkles")
}

// Add a sparkle!
func addSparkle(w http.ResponseWriter, h *http.Request) {
	var s Sparkle
	b := json.NewDecoder(h.Body)
	b.Decode(&s)

	result := db.AddSparkle(s)
	returnJSON(result, w, h)
}

// Get the entire data set
func getSparkles(w http.ResponseWriter, h *http.Request) {
	returnJSON(db.Sparkles, w, h)
}

// This only goes back 60 days
func topGiven(w http.ResponseWriter, h *http.Request) {
	afterDate := time.Now().AddDate(0, 0, -60)
	result := db.TopGiven(afterDate)
	returnJSON(result, w, h)
}

// This only goes back 60 days
func topReceived(w http.ResponseWriter, h *http.Request) {
	afterDate := time.Now().AddDate(0, 0, -60)
	result := db.TopReceived(afterDate)
	returnJSON(result, w, h)
}

// Get all the sparkles for someone in particular
func getSparklesForRecipient(w http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	rcpt := vars["recipient"]
	sparkles := db.SparklesForUser(rcpt)
	returnJSON(sparkles, w, h)
}

// Get stats for a user
func getStatsForUser(w http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	user := vars["user"]
	stats := StatsForUser(user)
	returnJSON(stats, w, h)
}

// Get the sparkle graph
func getSparkleGraph(w http.ResponseWriter, h *http.Request) {
	result := db.Graph()
	returnJSON(result, w, h)
}

// Migrate sparkles from one user to another
func migrateSparkles(w http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	from := vars["from"]
	to := vars["to"]
	db.MigrateSparkles(from, to)
	updatedSparkles := db.SparklesForUser(to)
	returnJSON(updatedSparkles, w, h)
}
