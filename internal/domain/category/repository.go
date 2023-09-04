package category

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"guhkun13/pizza-api/database"
)

type Repository struct {
	DB *database.Database
}

func NewRepository(db *database.Database) Repository {
	return Repository{
		DB: db,
	}
}

func (r *Repository) GetCategoryById(id int) (model Category, err error) {
	log.Info().Int("id", id).Msg("Repo.GetCategoryById")

	query := "SELECT * FROM category"
	log.Info().Str("query", query).Msg("check query")

	err = r.DB.PgxPool.Conn.Ping(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("error ping db")
	}

	rows, err := r.DB.PgxPool.Conn.Query(context.Background(), query)
	if err != nil {
		log.Error().Err(err).Msg("error query to db")
	}

	if rows.Err() == pgx.ErrNoRows {
		fmt.Println("NAH INI")
		fmt.Println(pgx.ErrNoRows.Error())
	}

	log.Info().
		Interface("rows", rows).
		Int64("affected", rows.CommandTag().RowsAffected()).
		Msg("response data from database")

	// err = row.Scan(model)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error on scan data from database")
	// 	return
	// }

	log.Info().Interface("model", model).Msg("OK")

	return
}
