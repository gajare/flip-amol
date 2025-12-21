package repository

import (
	"context"
	"fmt"
	"user-serice/internal/user/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
	UpdateUser(ctx context.Context, id int64, req *model.UpdateUserRequest) (*model.User, error)
	DeleteUser(ctx context.Context, id int64) (bool, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error) {
	query := `INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	var user model.User
	row := r.db.QueryRow(ctx, query, req.UserName, req.Email, req.Password, req.Role)
	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return &user, nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	query := `SELECT id, username, email, password, role, created_at FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	query := `SELECT id, username, email, password, role, created_at FROM users WHERE id = $1`
	var user model.User
	row := r.db.QueryRow(ctx, query, id)
	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int64, req *model.UpdateUserRequest) (*model.User, error) {
	query := `UPDATE users SET username = $1, email = $2, password = $3, role = $4 WHERE id = $5 RETURNING id, username, email, password, role, created_at`
	var user model.User
	row := r.db.QueryRow(ctx, query, req.UserName, req.Email, req.Password, req.Role, id)
	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	return &user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int64) (bool, error) {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete user: %v", err)
	}
	return true, nil
}
