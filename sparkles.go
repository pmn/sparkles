package main

import (
	"fmt"
	"net/http"
	"time"
)

type Request struct {
	SparkledBy string    `json:"sparkled_by"`
	Recipient  string    `json:"recipient"`
	SparkledAt time.Time `json:"sparkled_at"`
}

func defaultHandler(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Default sparkles")
}

func addSparkles(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Add a sparkle")
}

func getSparkles(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Get sparkles")
}
