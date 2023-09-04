package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pseidemann/finish"
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
	"guhkun13/pizza-api/database"
	"guhkun13/pizza-api/internal/domain/category"
	"guhkun13/pizza-api/internal/domain/product"
	"guhkun13/pizza-api/router"
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
	db := database.NewDatabase(&env)
	defer db.Close()

	// setup handler
	categoryRepo := category.NewRepository(db)
	category.NewService(&env, categoryRepo)

	productRepo := product.NewRepository(db)
	srv := product.NewService(&env, productRepo)
	productHandler := product.NewHandler(&env, *srv)

	domainHandlers := router.DomainHandlers{
		Product: productHandler,
	}

	// root router
	router := router.NewRootRouter(domainHandlers)

	// start server
	server := Server{
		Port:   3333,
		Router: router.Init(),
		DB:     db,
	}
	server.Start2()
}

type Server struct {
	Port   int
	Router http.Handler
	DB     *database.Database
}

func (s Server) Start2() {
	addr := fmt.Sprintf(":%d", s.Port)
	http.ListenAndServe(addr, s.Router)
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

	if err != http.ErrServerClosed {
		log.Fatal().Err(err)
	}

	fin.Wait()
}
