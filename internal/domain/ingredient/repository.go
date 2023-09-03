package ingredient

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/database"
)

type Repository struct {
	db *database.PgSqlConn
}

func NewRepository(dbConn *database.PgSqlConn) Repository {
	log.Info().Msg("NewRepository")
	return Repository{
		db: dbConn,
	}
}

func (c Repository) GetIngredients() ([]Ingredient, error) {
	log.Info().Msg("Repository.GetIngredients")
	rows, err := c.db.Conn.Query(context.Background(), "SELECT * from ingredient")
	if err != nil {
		log.Error().Err(err).Msg("something error")
		os.Exit(1)
	}

	rowsAffected := rows.CommandTag().RowsAffected()
	log.Info().Int("aff", int(rowsAffected)).Msg("affected")

	if rows.Next() {
		log.Info().Msg("has next!!!!!!!!!!!!!")
		vals, _ := rows.Values()

		log.Info().
			Interface("vals", vals).
			Msg("vals pools")

		// yang ini gagal!
		var ing Ingredient

		fmt.Println("mau nge-scan")
		rows.Scan(&ing)
		log.Info().
			Interface("ing", ing.ID).
			Msg("ing val -----")
	}
	log.Info().
		Interface("rows", rows).
		Msg("rows from db pool here")

	log.Info().Msg("CARA KEDUAAA")

	rows2, err := c.db.Conn.Query(context.Background(), "SELECT * from ingredient")
	if err != nil {
		log.Error().Err(err).Msg("something error")
		os.Exit(1)
	}

	// this works!
	ingredients, err := pgx.CollectRows(rows2, pgx.RowToStructByName[Ingredient])
	log.Info().Interface("result", ingredients).Msg("Collect rows")

	return ingredients, nil
}
