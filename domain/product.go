package domain

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/guregu/null"
)

type (
	Product struct {
		ID   int         `db:"id",json:"id"`
		Name null.String `db:"name",json:"name,omitempty"`
	}

	// Repository Interface
	ProductRepository interface {
		GetProductById(ctx context.Context, id int) Product
	}

	// Service Interface
	ProductService interface {
		GetProductById(ctx context.Context, id int) Product
	}

	// Handler Interface
	ProductHandler interface {
		GetProductById() chi.Router
	}
)
