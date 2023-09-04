package category

import "guhkun13/pizza-api/database"

type Repository struct {
	DB *database.PgSqlConn
}

func NewRepository(db *database.PgSqlConn) Repository {
	return Repository{
		DB: db,
	}
}
