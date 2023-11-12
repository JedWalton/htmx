package main

import (
	"fmt"
	"htmx/film"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set.")
	}
	fmt.Printf("Listening on port %v\n", port)

	db, err := NewPostgreSQL()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Init packages here
	film.InitFilm(db)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
