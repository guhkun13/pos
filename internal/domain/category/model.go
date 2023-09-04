package category

import "github.com/guregu/null"

type Category struct {
	ID   int64       `db:"id"`
	Name null.String `db:"name"`
}
