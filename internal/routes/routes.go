package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", indexHandler)
	r.Get("/api/data", apiDataHandler)

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to homepage")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "some data from API")
}
