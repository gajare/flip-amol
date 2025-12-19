package repository

import (
	"context"
	"fmt"
	"product-service/internal/product/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) (*model.Product, error)
	GetAll(ctx context.Context) ([]*model.Product, error)
	GetByID(ctx context.Context, id string) (*model.Product, error)
	Update(ctx context.Context, id string, update *model.UpdateProductRequest) (*model.Product, error)
	Delete(ctx context.Context, id string) error
}

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	query := `
		INSERT INTO products (id, name, description, price, stock, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, name, description, price, stock, created_at, updated_at`

	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	err := r.db.QueryRow(ctx, query,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.CreatedAt,
		product.UpdatedAt,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return product, nil
}

func (r *productRepository) GetAll(ctx context.Context) ([]*model.Product, error) {
	query := `
		SELECT id, name, description, price, stock, created_at, updated_at
		FROM products
		ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, &product)
	}

	return products, nil
}

func (r *productRepository) GetByID(ctx context.Context, id string) (*model.Product, error) {
	query := `
		SELECT id, name, description, price, stock, created_at, updated_at
		FROM products
		WHERE id = $1`

	var product model.Product
	err := r.db.QueryRow(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get product by id: %w", err)
	}

	return &product, nil
}

func (r *productRepository) Update(ctx context.Context, id string, update *model.UpdateProductRequest) (*model.Product, error) {
	query := `
		UPDATE products
		SET 
			name = COALESCE($2, name),
			description = COALESCE($3, description),
			price = COALESCE($4, price),
			stock = COALESCE($5, stock),
			updated_at = $6
		WHERE id = $1
		RETURNING id, name, description, price, stock, created_at, updated_at`

	now := time.Now()
	var product model.Product
	err := r.db.QueryRow(ctx, query,
		id,
		update.Name,
		update.Description,
		update.Price,
		update.Stock,
		now,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return &product, nil
}

func (r *productRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}
