package main

import "net/http"

func initRoutes() {
	// define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)
}
