package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
)

type PgxPoolConn struct {
	Conn *pgxpool.Pool
}

func NewPgxPoolConn(env *config.EnvironmentVariables) *PgxPoolConn {
	return &PgxPoolConn{
		Conn: CreateConnection(env),
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

	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Error().Err(err).Msgf("unable to connect to database %s", dsn)
		panic("unable connect database")
	}

	err = conn.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	log.Info().Msg("database is connected")

	return conn
}
