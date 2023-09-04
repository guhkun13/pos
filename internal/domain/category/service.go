package category

import (
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
)

type Service struct {
	// Repo Repository
	Env *config.EnvironmentVariables
}

func NewService(env *config.EnvironmentVariables) *Service {
	return &Service{
		// Repo: repo,
		Env: env,
	}
}

func (s *Service) GetCategory(id int) (res CategoryResponse) {
	log.Info().Int("id", id).Msg("GetCategory")
	// // category, err := s.Repo.GetProductById(id)
	// if err != nil {
	// 	log.Error().Err(err).Msg("error after call repo")
	// }

	// log.Info().Interface("res", res).Msg("before")
	// res = res.New(category)
	// log.Info().Interface("res", res).Msg("after")

	return
}
