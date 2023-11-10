package main

import (
	"database/sql"
	"fmt"
	"htmx/film"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Listening on port 8000...")

	db, err := sql.Open("postgres", os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	film.InitFilm(db)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
