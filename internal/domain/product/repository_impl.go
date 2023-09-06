package product

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func (r Repository) GetProductById(ctx context.Context, id int) (product Product) {
	log.Info().Int("id", id).Msg("Repository.GetProductById")
	sql := fmt.Sprintf("select * from product where id=%d", id)
	log.Info().Msgf("sql = %s ", sql)

	row := r.DB.PgxPool.Conn.QueryRow(ctx, sql)
	var err error
	if row.Scan(product); err != nil {
		log.Error().Err(err).Msg("failed get product")
	}

	rows, err := r.DB.PgxPool.Conn.Query(ctx, sql)
	if err != nil {
		log.Error().Err(err).Msg("failed get product")
	}

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[Product])
	log.Info().Interface("result", items).Msg("Collect rows")
	if err != nil {
		fmt.Println("ERROR WHEN SCAN")
		log.Error().Err(err).Msg("EM")
	}
	product = items[0]
	log.Info().Interface("product", product).Msg("final result")
	return
}
