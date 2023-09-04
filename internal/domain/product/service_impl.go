package product

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (s *Service) GetProduct(ctx context.Context, id int) string {
	log.Info().Msg("Service GetProduct")

	name := "truck"
	category := "cars"
	// category := s.CategoryService.GetCategory(id)
	// log.Info().Interface("category", category).Msg("result category")

	product := s.Repo.GetProductById(ctx, id)
	log.Info().Interface("product", product).Msg("result product")
	name = product.Name.String

	res := fmt.Sprintf("nama product: %s dari kategori: %s", name, category)
	return res
}

func (s *Service) CreateProduct(ctx context.Context, name string) string {
	log.Info().Msg("Service CreateProduct")

	return "yes"
}
