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

func defaultHandler(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Default sparkles")
}

func addSparkles(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Add a sparkle")
	var s Sparkle
	b := json.NewDecoder(h.Body)
	b.Decode(&s)

	db.AddSparkle(s)
	fmt.Printf("%v", db.Sparkles)
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
