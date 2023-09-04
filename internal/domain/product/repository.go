package product

import (
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

// type RepositoryIface interface {
// 	GetProductById(ctx context.Context, id int) Product
// }
