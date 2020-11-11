package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var config Config

func main() {
	config = GetConfig()

	log.Println("Starting server...")
	r := mux.NewRouter()
	r.HandleFunc("/entries/", JSONHandler)
	r.HandleFunc("/feed/", RSSHandler)
	r.HandleFunc("/atom.xml", AtomHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Your feed:")
	log.Printf("\tJSON: %s:8000/entries/", config.ServerHostname)
	log.Printf("\tRSS: %s:8000/feed/", config.ServerHostname)
	log.Printf("\tAtom: %s:8000/atom.xml", config.ServerHostname)

	log.Fatal(srv.ListenAndServe())
}
