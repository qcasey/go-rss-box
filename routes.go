package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONHandler to debug entries with JSON
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	feed, err := CreateFeed()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(fmt.Sprintf("Failed to parse entries: %s", err.Error()))
		return
	}

	rss, err := feed.ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintln(w, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, rss)
}

// RSSHandler fetches recent mail and responds with an RSS feed
func RSSHandler(w http.ResponseWriter, r *http.Request) {
	feed, err := CreateFeed()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(fmt.Sprintf("Failed to parse entries: %s", err.Error()))
		return
	}

	rss, err := feed.ToRss()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintln(w, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, rss)
}

// AtomHandler fetches recent mail and responds with an Atom feed
func AtomHandler(w http.ResponseWriter, r *http.Request) {
	feed, err := CreateFeed()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(fmt.Sprintf("Failed to parse entries: %s", err.Error()))
		return
	}

	rss, err := feed.ToAtom()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintln(w, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, rss)
}
