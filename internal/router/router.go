package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type DomainHandler struct {
}
type Router struct {
	DomainHandler DomainHandler
}

func InitRouter() http.Handler {
	r := chi.NewRouter()

	setupMiddleware(r)
	setupRoutes(r)

	return r
}

func setupMiddleware(mux *chi.Mux) {
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
}

func setupRoutes(mux *chi.Mux) {
	mux.Route("/v1", func(r chi.Router) {
		v1Routes(r)
	})
	mux.Route("/v2", func(r chi.Router) {
		v2Routes(r)
	})
}

func v1Routes(r chi.Router) chi.Routes {
	r.Get("/", indexHandler)
	r.Get("/api/data", apiDataHandler)

	return r
}

func v2Routes(r chi.Router) chi.Routes {
	r.Get("/", indexHandlerV2)
	r.Get("/api/data", apiDataHandlerV2)

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to homepage")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "some data from API")
}

func indexHandlerV2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to homepage v2")
}

func apiDataHandlerV2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "some data from API v2")
}
