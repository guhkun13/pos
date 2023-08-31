package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
	"guhkun13/pizza-api/internal/routes"
)

func main() {
	config.SetupLogger()

	env, errEnv := config.LoadEnv()
	if errEnv != nil {
		log.Error().Err(errEnv)
	}

	log.Info().Str("db host", env.Database.Host).Msg("config")
	log.Info().Str("db host", env.Database.Password).Msg("config")

	router := routes.NewRouter()
	port := 3333
	addr := fmt.Sprintf(":%d", port)
	log.Info().Msgf("Start server at %s", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}

}
