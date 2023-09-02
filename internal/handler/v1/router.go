package v1

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
	"guhkun13/pizza-api/database"
	"guhkun13/pizza-api/internal/domain/ingredient"
)

func SetupRouter(r chi.Router) chi.Routes {
	r.Get("/", indexHandler)
	r.Route("/ingredient", func(r chi.Router) {
		IngredientRouter(r)
	})

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to homepage")
}

// Ingredient
func IngredientRouter(r chi.Router) chi.Routes {
	r.Get("/", ingredientHandler)
	return r
}
func ingredientHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("ingredientHandler")

	env, err := config.LoadEnv()
	if err != nil {
		log.Error().Err(err)
	}
	log.Info().Msg("LoadEnv OK")

	db := database.NewPgSqlConnection(&env)
	repo := ingredient.NewRepository(db)

	service := ingredient.NewIngredientService(&env, repo)
	res := service.Get()

	fmt.Fprintln(w, res)
}
