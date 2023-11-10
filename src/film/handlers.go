package film

import (
	"html/template"
	"net/http"
	"time"
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
	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	tmpl.Execute(w, films)
}

// handler function #2 - returns the template block with the newly added film, as an HTMX response
func (h *Handlers) H2(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
	// tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl := template.Must(template.ParseFiles("fe/index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}
