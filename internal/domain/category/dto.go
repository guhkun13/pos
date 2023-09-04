package category

import "github.com/guregu/null"

type CategoryResponse struct {
	ID   int64       `json:"id"`
	Name null.String `json:"name"`
}

func (dto *CategoryResponse) New(data Category) CategoryResponse {
	return CategoryResponse{
		ID:   data.ID,
		Name: data.Name,
	}
}
