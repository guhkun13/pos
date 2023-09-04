package product

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
)

type Handler struct {
	Env     *config.EnvironmentVariables
	Service *Service
}

func NewHandler(env *config.EnvironmentVariables, service *Service) *Handler {
	return &Handler{
		Env:     env,
		Service: service,
	}
}

func (h *Handler) Handlers(rc chi.Router) {
	rc.Get("/product", h.GetProductHandler)
	rc.Route("/sub", func(r chi.Router) {
		h.SubProductHandler(r)
	})
}

func (h *Handler) SubProductHandler(rc chi.Router) {
	log.Info().Msg("SubProductHandler")
	rc.Get("/x", h.GetSubProductHandler)
}

func (h *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	res := h.Service.GetProduct()
	w.Write([]byte(res))
}

func (h *Handler) GetSubProductHandler(w http.ResponseWriter, r *http.Request) {
	res := h.Service.GetProduct()
	w.Write([]byte(res))
}
