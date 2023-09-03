package product

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"guhkun13/pizza-api/config"
)

type Handler struct {
	Env *config.EnvironmentVariables
}

func NewHandler(env *config.EnvironmentVariables) Handler {
	return Handler{
		Env: env,
	}
}

func (h Handler) Handlers(rc chi.Router) {
	rc.Get("/product", h.GetProductHandler)
	h.SubProductHandler(rc)
}

func (h Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get product"))
}

func (h Handler) GetSubProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get sub-product"))
}

func (h *Handler) SubProductHandler(rc chi.Router) {

	rc.Get("/sub-product", h.GetSubProductHandler)
}
