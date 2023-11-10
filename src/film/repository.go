package film

import (
	"database/sql"
	"log"
)

// FilmRepository defines the interface for film data operations.
type FilmRepository interface {
	GetAllFilms() ([]Film, error)
	AddFilm(film Film) error
	// ... other CRUD operations
}

type filmRepository struct {
	db *sql.DB
}

func NewFilmRepository(db *sql.DB) FilmRepository {
	return &filmRepository{db: db}
}

func (repo *filmRepository) GetAllFilms() ([]Film, error) {
	var films []Film

	rows, err := repo.db.Query("SELECT title, director FROM films")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var f Film
		if err := rows.Scan(&f.Title, &f.Director); err != nil {
			return nil, err
		}
		films = append(films, f)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	log.Printf("films: %v", films)
	return films, nil
}

func (repo *filmRepository) AddFilm(film Film) error {
	// SQL statement to insert a new film
	sqlStatement := `INSERT INTO films (title, director) VALUES ($1, $2)`

	// Executing the SQL statement
	_, err := repo.db.Exec(sqlStatement, film.Title, film.Director)
	if err != nil {
		// Handle any errors that occur during the insert
		return err
	}

	// Return nil if no errors occurred
	return nil
}
