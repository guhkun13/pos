package product

import (
	"context"

	"guhkun13/pizza-api/config"
)

type Service struct {
	Env  *config.EnvironmentVariables
	Repo Repository
}

func NewService(
	env *config.EnvironmentVariables,
	repo Repository,
) Service {
	return Service{
		Env:  env,
		Repo: repo,
	}
}

// Interface sbg kontrak, apa saja method2 yang mesti di-implement oleh Service
type ServiceInterface interface {
	GetProduct(ctx context.Context, id int) string
	CreateProduct(ctx context.Context, name string) string
}
