package film

import (
	"database/sql"
	"net/http"
)

func InitFilm(db *sql.DB) {
	filmRepo := NewFilmRepository(db)
	filmHandlers := NewHandlers(filmRepo)

	http.HandleFunc("/", filmHandlers.H1)
	http.HandleFunc("/add-film/", filmHandlers.H2)
}
