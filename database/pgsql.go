package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/config"
	"guhkun13/pizza-api/internal/domain/category"
)

type PgSqlConn struct {
	Conn *pgxpool.Pool
}

func NewPgSqlConn(env *config.EnvironmentVariables) *PgSqlConn {
	return &PgSqlConn{
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
		panic("gagal konek")
	}

	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Println("PANIC BANG")
		panic(err)
	}

	log.Info().Msg("database is connected")
	TestQuery(conn)

	return conn
}

func TestQuery(conn *pgxpool.Pool) {
	fmt.Println("TEST QUERY BANG")
	rows, err := conn.Query(context.Background(), "SELECT * from category")
	if err != nil {
		log.Error().Err(err).Msg("something error")
	}

	// var items []category.Category

	fmt.Println("CEK NEXT")
	if rows.Next() {
		fmt.Println("ADA NEXT!")
		vals, err := rows.Values()
		if err != nil {
			fmt.Println("NAH ERROR")
		}

		items, err := pgx.CollectRows(rows, pgx.RowToStructByName[category.Category])
		log.Info().Interface("result", items).Msg("Collect rows")
		if err != nil {
			fmt.Println("ERROR WHEN SCAN")
			log.Error().Err(err).Msg("EM")
		}

		log.Info().
			Interface("vals", vals).
			Interface("scanned", items).
			Msg("result")
	}

	fmt.Println("DONE")

}
