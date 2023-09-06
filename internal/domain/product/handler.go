package product

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	response "guhkun13/pizza-api/shared"
)

type Handler struct {
	// Env     *config.EnvironmentVariables
	Service Service
}

func NewHandler(
	// env *config.EnvironmentVariables,
	srv Service) Handler {
	return Handler{
		// Env:     env,
		Service: srv,
	}
}

func (h *Handler) Handlers(mux chi.Router) {
	mux.Route("/product", func(r chi.Router) {
		r.Get("/{id}", h.GetProduct)
		r.Get("/html", h.ShowHtml)
	})

}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func (h *Handler) ShowHtml(w http.ResponseWriter, r *http.Request) {

	tmplf, err := template.ParseFiles("html/layout.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl := template.Must(tmplf, err)

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task A", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
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
