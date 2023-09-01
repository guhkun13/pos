package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
	"guhkun13/pizza-api/internal/router"
)

var env *config.EnvironmentVariables

// setups: logger, config, router + middleware, setup docs, graceful shutdown, run!
func main() {
	// init logger
	config.InitLogger()

	_, err := config.LoadEnv()
	if err != nil {
		log.Error().Err(err).Msg("load env error")
	}

	router := router.InitRouter()
	port := 3333
	addr := fmt.Sprintf(":%d", port)
	log.Info().Msgf("Start server at %s", addr)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}

}
