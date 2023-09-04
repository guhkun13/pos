package product

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
	"guhkun13/pizza-api/internal/domain/category"
)

type Service struct {
	Env             *config.EnvironmentVariables
	CategoryService *category.Service
}

func NewService(
	env *config.EnvironmentVariables,
	categorySrv *category.Service,
) *Service {
	return &Service{
		Env:             env,
		CategoryService: categorySrv,
	}
}

func (s *Service) GetProduct() string {
	log.Info().Msg("Service GetProduct")

	name := "truck"
	// category := "toys"
	category := s.CategoryService.GetCategory(1)
	log.Info().Interface("category", category).Msg("in service")

	res := fmt.Sprintf("nama product: %s dari kategori: %s", name, category)
	return res
}
