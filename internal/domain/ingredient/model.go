package ingredient

import (
	"github.com/guregu/null"
)

type Ingredient struct {
	ID          int         `db:"id"`
	Name        null.String `db:"name"`
	Description null.String `db:"description"`
	Thumbnail   null.String `db:"thumbnail"`
	CreatedAt   null.Time   `db:"created_at"`
	UpdatedAt   null.Time   `db:"updated_at"`
	CreatedBy   null.Int    `db:"created_by"`
	UpdatedBy   null.Int    `db:"updated_by"`
}
