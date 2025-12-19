package service

import (
	"context"
	"product-service/internal/product/model"
	"product-service/internal/product/repository"

	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(ctx context.Context, input model.CreateProductRequest) (*model.Product, error)
	GetProducts(ctx context.Context) ([]*model.Product, error)
	GetProductByID(ctx context.Context, id string) (*model.Product, error)
	UpdateProduct(ctx context.Context, id string, input model.UpdateProductRequest) (*model.Product, error)
	DeleteProduct(ctx context.Context, id string) (bool, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(ctx context.Context, input model.CreateProductRequest) (*model.Product, error) {
	product := &model.Product{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}

	return s.repo.Create(ctx, product)
}

func (s *productService) GetProducts(ctx context.Context) ([]*model.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *productService) GetProductByID(ctx context.Context, id string) (*model.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *productService) UpdateProduct(ctx context.Context, id string, input model.UpdateProductRequest) (*model.Product, error) {
	return s.repo.Update(ctx, id, &input)
}

func (s *productService) DeleteProduct(ctx context.Context, id string) (bool, error) {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}