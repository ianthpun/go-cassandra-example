package main

import (
	"log"
	"net/http"

	"github.com/ianthpun/go-cassandra-example/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Heartbeat)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
