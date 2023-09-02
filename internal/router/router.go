package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	v1 "guhkun13/pizza-api/internal/handler/v1"
)

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
		v1.SetupRouter(r)
	})
}
