package main

import (
	"database/sql"
	"fmt"
	"htmx/film"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq" // This is the PostgreSQL driver for database/sql
)

// Define a struct to hold the db connection
type AppHandler struct {
	db *sql.DB
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set.")
	}

	var db *sql.DB
	var err error

	// Retry logic for database connection
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = NewPostgreSQL()
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database: %v, retrying in %v seconds\n", err, 5)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to connect to database after %d retries: %v", maxRetries, err)
	}
	defer db.Close()

	handler := &AppHandler{db: db}
	http.HandleFunc("/health", handler.HealthCheckHandler)

	// Init packages here
	film.InitFilm(db)

	fmt.Printf("Listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (h *AppHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Use h.db for database operations
	if err := h.db.Ping(); err != nil {
		// If the database is not reachable, return an error response
		http.Error(w, "Database not reachable", http.StatusInternalServerError)
		return
	}

	// You can add more checks here if necessary

	// If everything is okay, send a success response
	fmt.Fprintln(w, "OK")
}
