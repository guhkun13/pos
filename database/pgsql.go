package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
)

type PgSqlConnection struct {
	PgPool *pgxpool.Pool
}

func NewPgSqlConnection(config *config.EnvironmentVariables) *PgSqlConnection {
	return &PgSqlConnection{
		PgPool: CreateConnection(config),
	}
}

func CreateConnection(config *config.EnvironmentVariables) *pgxpool.Pool {
	log.Info().Msg("CreateConnection")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	log.Info().Str("dsn", dsn).Msg("dsn")

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Error().Err(err)
		os.Exit(1)
	}

	log.Info().Msg("return dbpool")

	return dbpool
}
