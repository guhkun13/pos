package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pseidemann/finish"
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
	"guhkun13/pizza-api/database"
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
	db := database.NewPgSqlConn(&env)
	defer db.Conn.Close()

	fmt.Println("HALO DISINI YAK")
	rows, err := db.Conn.Query(context.Background(), "select * from ingredient")
	if err != nil {
		fmt.Println("ERORR CUK")
	}
	log.Info().
		Interface("rows", rows).
		Interface("rows aff", rows.CommandTag().RowsAffected()).
		Msg("result")

	// setup handler
	// categoryRepo := category.NewRepository(db)
	// categorySrv := category.NewService(&env, categoryRepo)

	// productService := product.NewService(&env, categorySrv)
	// productHandler := product.NewHandler(&env, productService)

	domainHandlers := DomainHandlers{
		// Product: productHandler,
	}

	// root router
	router := NewRootRouter(domainHandlers)

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
	DB     *database.PgSqlConn
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

	defer s.DB.Conn.Close()

	fin.Wait()
}
