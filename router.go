package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"guhkun13/pizza-api/internal/domain/product"
)

type RootRouter struct {
	DomainHandlers DomainHandlers
}

func NewRootRouter(handlers DomainHandlers) RootRouter {
	return RootRouter{
		DomainHandlers: handlers,
	}
}

type DomainHandlers struct {
	Product *product.Handler
}

func (root *RootRouter) Init() http.Handler {
	mux := chi.NewRouter()

	root.SetupMiddleware(mux)
	root.SetupRoutesV1(mux)

	return mux
}

func (root *RootRouter) SetupMiddleware(mux *chi.Mux) {
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
}

func (root *RootRouter) SetupRoutesV1(mux *chi.Mux) {
	mux.Route("/", func(r chi.Router) {
		root.DomainHandlers.Product.Handlers(r)
	})
}
