package film

import (
	"database/sql"
	"net/http"
)

func initFilmHandlers(db *sql.DB) *Handlers {
	filmRepo := NewFilmRepository(db)
	return NewHandlers(filmRepo)
}

func InitFilm(db *sql.DB) {
	filmHandlers := initFilmHandlers(db)

	http.HandleFunc("/", filmHandlers.H1)
	http.HandleFunc("/add-film/", filmHandlers.H2)
}
