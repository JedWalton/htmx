package main

import (
	"fmt"
	"htmx/film"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening on port 8000...")

	db, err := NewPostgreSQL()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Init packages here
	film.InitFilm(db)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
