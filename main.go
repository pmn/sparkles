package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
  "os"
)

type Request struct {
	SparkledBy string    `json:"sparkled_by"`
	Recipient  string    `json:"recipient"`
	SparkledAt time.Time `json:"sparkled_at"`
}

func defaultHandler(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Default sparkles")
}

func addSparkles(w http.ResponseWriter, h *http.Request){
  fmt.Fprint(w, "Add a sparkle")
}

func getSparkles(w http.ResponseWriter, h *http.Request){
  fmt.Fprint(w, "Get sparkles")
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteAddr := r.RemoteAddr
		if len(remoteAddr) == 0 {
			remoteAddr = r.Header.Get("x-forwarded-for")
		}
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler).Methods("GET")
  r.HandleFunc("/sparkles", addSparkles).Methods("POST")
  r.HandleFunc("/sparkles", getSparkles).Methods("GET")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.Handle("/", r)
	log.Printf("[+] Starting the sparkles app on port %s", port)
	http.ListenAndServe(":"+port, Log(http.DefaultServeMux))
}
