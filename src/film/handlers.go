package film

import (
	"html/template"
	"net/http"
)

type Handlers struct {
	Repo FilmRepository
}

func NewHandlers(repo FilmRepository) *Handlers {
	return &Handlers{Repo: repo}
}

// handler function #1 - returns the index.html template, with film data
func (h *Handlers) H1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("fe/index.html"))

	// Use GetAllFilms method to fetch films from the database
	films, err := h.Repo.GetAllFilms()
	if err != nil {
		// Handle error appropriately, e.g., send an HTTP error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a data structure to pass to the template
	data := map[string][]Film{
		"Films": films,
	}

	// Execute the template with the fetched films
	if err := tmpl.Execute(w, data); err != nil {
		// Handle template execution error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// handler function #2 - adds a new film to the database and returns the template block with the newly added film, as an HTMX response
func (h *Handlers) H2(w http.ResponseWriter, r *http.Request) {
	// Parse form values
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	director := r.FormValue("director")

	// Create a new Film instance
	newFilm := Film{Title: title, Director: director}

	// Add the new film to the database
	if err := h.Repo.AddFilm(newFilm); err != nil {
		http.Error(w, "Failed to add film", http.StatusInternalServerError)
		return
	}

	// Load the template
	tmpl := template.Must(template.ParseFiles("fe/index.html"))

	// Execute the specific template block for the newly added film
	if err := tmpl.ExecuteTemplate(w, "film-list-element", newFilm); err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}
