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
		log.Fatalf("Failed to connect to database: %v", err)
	}

	film.InitRoutes()

	log.Fatal(http.ListenAndServe(":8000", nil))
}
