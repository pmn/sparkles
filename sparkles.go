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

var sparkles []Sparkle

func defaultHandler(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Default sparkles")
}

func addSparkles(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Add a sparkle")
	var s Sparkle
	b := json.NewDecoder(h.Body)
	b.Decode(&s)

	sparkles = append(sparkles, s)
	fmt.Printf("%v", sparkles)
}

func getSparkles(w http.ResponseWriter, h *http.Request) {
	fmt.Fprintf(w, "%v", sparkles)
}

func getSparklesForRecipient(w http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	rcpt := vars["recipient"]
	fmt.Fprint(w, "Get sparkles for ", rcpt)
}
