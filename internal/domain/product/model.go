package product

import "github.com/guregu/null"

type Product struct {
	ID   int         `db:"id"`
	Name null.String `db:"name"`
}
