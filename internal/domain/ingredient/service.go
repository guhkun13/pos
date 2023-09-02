package ingredient

import (
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
)

type IngredientService struct {
	Config     *config.EnvironmentVariables
	Repository Repository
}

func NewIngredientService(
	config *config.EnvironmentVariables,
	repo Repository) *IngredientService {
	return &IngredientService{
		Config:     config,
		Repository: repo,
	}
}

func (s IngredientService) Get() string {
	log.Info().Msg("IngredientService.Get")

	rows, err := s.Repository.GetIngredients()
	log.Info().Msg("from service call Repository.GetIngredients")

	if err != nil {
		log.Error().Err(err).Msg("ERORR ")
	}
	log.Info().Interface("rows", rows).Msg("interface rows")

	return "get ingredient"
}
