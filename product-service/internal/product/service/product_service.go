package service

import (
	"context"

	"product-service/internal/product/model"
	"product-service/internal/product/repository"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req model.CreateProductRequest) (*model.Product, error)
	GetProducts(ctx context.Context) ([]*model.Product, error)
	GetProductByID(ctx context.Context, id string) (*model.Product, error)
	DeleteProduct(ctx context.Context, id string) (bool, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(
	ctx context.Context,
	req model.CreateProductRequest,
) (*model.Product, error) {
	return s.repo.Create(ctx, req)
}

func (s *productService) GetProducts(
	ctx context.Context,
) ([]*model.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *productService) GetProductByID(
	ctx context.Context,
	id string,
) (*model.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *productService) DeleteProduct(
	ctx context.Context,
	id string,
) (bool, error) {
	return s.repo.Delete(ctx, id)
}
