package category

import "guhkun13/pizza-api/config"

type Service struct {
	Repo Repository
	Env  *config.EnvironmentVariables
}

func NewService(env *config.EnvironmentVariables, repo Repository) *Service {
	return &Service{
		Repo: repo,
		Env:  env,
	}
}
