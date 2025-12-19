package graph

import (
	"context"
	"product-service/graph/model"
	"product-service/internal/product/service"
)

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

type Resolver struct {
	productService service.ProductService
}

func NewResolver(productService service.ProductService) *Resolver {
	return &Resolver{productService: productService}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return r
}

func (r *Resolver) Query() generated.QueryResolver {
	return r
}

func (r *Resolver) CreateProduct(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	req := service.CreateProductRequest{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}

	product, err := r.productService.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (r *Resolver) UpdateProduct(ctx context.Context, id string, input model.UpdateProductInput) (*model.Product, error) {
	req := service.UpdateProductRequest{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}

	product, err := r.productService.UpdateProduct(ctx, id, req)
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (r *Resolver) DeleteProduct(ctx context.Context, id string) (bool, error) {
	return r.productService.DeleteProduct(ctx, id)
}

func (r *Resolver) Products(ctx context.Context) ([]*model.Product, error) {
	products, err := r.productService.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	var result []*model.Product
	for _, p := range products {
		result = append(result, &model.Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			CreatedAt:   p.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   p.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return result, nil
}

func (r *Resolver) Product(ctx context.Context, id string) (*model.Product, error) {
	product, err := r.productService.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}