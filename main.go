package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Log what's up
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteAddr := r.RemoteAddr
		if len(remoteAddr) == 0 {
			remoteAddr = r.Header.Get("x-forwarded-for")
		}
		log.Printf("%s %s %s", remoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

var db SparkleDatabase

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sparkles", addSparkle).Methods("POST")
	r.HandleFunc("/sparkles", getSparkles).Methods("GET")
	r.HandleFunc("/top/giver", topGiven).Methods("GET")
	r.HandleFunc("/top/receiver", topReceived).Methods("GET")
	r.HandleFunc("/sparkles/{recipient}", getSparklesForRecipient).Methods("GET")
	adminMode := os.Getenv("SPARKLE_ADMIN_MODE")
	if adminMode == "TRUE" {
		r.HandleFunc("/migrate/{from}/{to}", migrateSparkles)
	}
	r.HandleFunc("/graph", getSparkleGraph).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	// Load the database from file
	db = LoadDB()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.Handle("/", r)
	log.Printf("[+] Starting the sparkles app on port %s", port)
	http.ListenAndServe(":"+port, Log(http.DefaultServeMux))
	log.Printf("... started!")
}
