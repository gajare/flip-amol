package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"product-service/internal/product/model"
)

type ProductRepository interface {
	Create(ctx context.Context, req model.CreateProductRequest) (*model.Product, error)
	GetAll(ctx context.Context) ([]*model.Product, error)
	GetByID(ctx context.Context, id string) (*model.Product, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(
	ctx context.Context,
	req model.CreateProductRequest,
) (*model.Product, error) {

	query := `
		INSERT INTO products (name, description, price, stock)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, description, price, stock, created_at, updated_at
	`

	row := r.db.QueryRow(ctx, query,
		req.Name,
		req.Description,
		req.Price,
		req.Stock,
	)

	var p model.Product
	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Stock,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepository) GetAll(ctx context.Context) ([]*model.Product, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, name, description, price, stock, created_at, updated_at
		FROM products
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*model.Product

	for rows.Next() {
		var p model.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return products, nil
}

func (r *productRepository) GetByID(
	ctx context.Context,
	id string,
) (*model.Product, error) {

	row := r.db.QueryRow(ctx, `
		SELECT id, name, description, price, stock, created_at, updated_at
		FROM products
		WHERE id = $1
	`, id)

	var p model.Product
	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Stock,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepository) Delete(
	ctx context.Context,
	id string,
) (bool, error) {

	cmd, err := r.db.Exec(ctx, `
		DELETE FROM products WHERE id = $1
	`, id)
	if err != nil {
		return false, err
	}

	return cmd.RowsAffected() > 0, nil
}
