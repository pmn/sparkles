package main

import (
  "fmt"
  "github.com/gorilla/mux"
  )

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", defaultHandler)
  r.HandleFunc("/leaders", leadersHandler)

  port := os.Getenv("PORT")
  if len(port) == 0 {
    port := "8080"
  }
}
