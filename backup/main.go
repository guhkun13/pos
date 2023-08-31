package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func loadEnv(key string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Error().Err(err).Msg("error on load env")
		return "", err
	}
	return os.Getenv(key), nil
}

func main() {
	fmt.Println("setup zerolog")
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	dbUrl, err := loadEnv("DATABASE_URL")
	log.Info().Interface("url", dbUrl).Msg("db details")
	if err != nil {
		log.Fatal().Msg("error connect database")
		return
	}

	dbpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		log.Err(err).Msg("Error on connect to db")
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
