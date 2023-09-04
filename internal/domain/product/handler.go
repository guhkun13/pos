package product

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
	response "guhkun13/pizza-api/shared"
)

type Handler struct {
	Env     *config.EnvironmentVariables
	Service Service
}

func NewHandler(env *config.EnvironmentVariables, srv Service) Handler {
	return Handler{
		Env:     env,
		Service: srv,
	}
}

func (h *Handler) Handlers(mux chi.Router) {
	mux.Route("/product", func(r chi.Router) {
		r.Get("/{id}", h.GetProduct)
		// r.Post("/", h.CreateProduct)
	})
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Handler.GetProduct")
	paramId := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		log.Error().Err(err).Msg("failed convert string to int")
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Int("id", id).Msg("id is valid")
	res := h.Service.GetProduct(r.Context(), id)

	// res := "dummy"
	response.ResponseWithJson(w, 200, res)
}

// func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	res := h.ServiceIface.CreateProduct(r.Context(), "gundam")
// 	w.Write([]byte(res))
// }
