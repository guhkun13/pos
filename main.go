package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pseidemann/finish"
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
	"guhkun13/pizza-api/database"
	"guhkun13/pizza-api/internal/domain/product/v1"
)

var env *config.EnvironmentVariables

// setups: logger, config, router + middleware, setup docs, graceful shutdown, run!
func main() {
	// init logger
	config.InitLogger()

	// load env
	env, err := config.LoadEnv()
	if err != nil {
		log.Error().Err(err).Msg("load env error")
		panic(err)
	}

	// init db
	db := database.NewPgSqlConnection(&env)

	// setup router
	// domain router
	productHandler := product.NewHandler(&env)

	// group of domain routers
	domainRouters := DomainHandlers{
		Product: productHandler,
	}

	// root router
	router := NewRootRouter(domainRouters)

	// start server
	server := Server{
		Port:   3333,
		Router: router.Init(),
		DB:     db,
	}
	server.Start()
}

type Server struct {
	Port   int
	Router http.Handler
	DB     *database.PgSqlConnection
}

func (s Server) Start() {
	addr := fmt.Sprintf(":%d", s.Port)
	log.Info().Msgf("Start server at %s", addr)
	srv := &http.Server{
		Addr:    addr,
		Handler: s.Router,
	}

	fin := finish.Finisher{
		Timeout: time.Second * 30,
	}
	err := srv.ListenAndServe()
	fin.Add(srv)

	go func() {

		if err != http.ErrServerClosed {
			log.Fatal().Err(err)
		}
	}()

	defer s.DB.PgPool.Close()

	fin.Wait()
}
