package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

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

var db SparkleDatabase

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler).Methods("GET")
	r.HandleFunc("/sparkles", addSparkles).Methods("POST")
	r.HandleFunc("/sparkles", getSparkles).Methods("GET")
	r.HandleFunc("/sparkles/{recipient}", getSparklesForRecipient).Methods("GET")

	// Load the database from file
	db.Load()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.Handle("/", r)
	log.Printf("[+] Starting the sparkles app on port %s", port)
	http.ListenAndServe(":"+port, Log(http.DefaultServeMux))
}
