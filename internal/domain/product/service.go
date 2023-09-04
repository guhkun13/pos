package product

import (
	"fmt"

	"guhkun13/pizza-api/config"
)

type Service struct {
	Env *config.EnvironmentVariables
}

func NewService(env *config.EnvironmentVariables) *Service {
	return &Service{
		Env: env,
	}
}

func (s *Service) GetProduct() string {

	name := "truck"
	category := "toys"

	return fmt.Sprintf("nama product: %s dari kategori: %s", name, category)
}
