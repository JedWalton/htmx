package film

import "net/http"

func InitRoutes() {
	// define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)
}
