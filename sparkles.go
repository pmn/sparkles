package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Sparkle struct {
	Sparkler string    `json:"sparkler"`
	Sparklee string    `json:"sparklee"`
	Reason   string    `json:"reason,omitempty"`
	Time     time.Time `json:"time,omitempty"`
}

type Leader struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// Return the data in JSON format. This is the default return method.
func returnJson(obj interface{}, w http.ResponseWriter, h *http.Request) {
	// Don't cache json returns. This is to work around ie's weird caching behavior
	w.Header().Set("Cache-Control", "no-cache")
	// Set the content type to json
	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprint(w, string(j))
}

func defaultHandler(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Default sparkles")
}

func addSparkles(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Add a sparkle")
	var s Sparkle
	b := json.NewDecoder(h.Body)
	b.Decode(&s)

	result := db.AddSparkle(s)
	returnJson(result, w, h)
}

func getSparkles(w http.ResponseWriter, h *http.Request) {
	fmt.Fprintf(w, "%v", db.Sparkles)
}

func getSparklesForRecipient(w http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	rcpt := vars["recipient"]
	fmt.Fprint(w, "Get sparkles for ", rcpt)
	sparkles := db.SparklesForUser(rcpt)
	fmt.Fprint(w, sparkles)
}
