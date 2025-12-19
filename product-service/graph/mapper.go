package graph

import (
	"time"

	graphModel "product-service/graph/model"
	internalModel "product-service/internal/product/model"
)

// -------- INPUT MAPPERS --------

func toServiceCreate(input graphModel.CreateProductInput) internalModel.CreateProductRequest {
	return internalModel.CreateProductRequest{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}
}

// -------- OUTPUT MAPPERS --------

func toGraphProduct(p *internalModel.Product) *graphModel.Product {
	if p == nil {
		return nil
	}

	return &graphModel.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		CreatedAt:   formatTime(p.CreatedAt),
		UpdatedAt:   formatTime(p.UpdatedAt),
	}
}

func formatTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
