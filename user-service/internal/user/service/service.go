package service

import (
	"context"
	"user-serice/internal/user/model"
	"user-serice/internal/user/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
	UpdateUser(ctx context.Context, id int64, req *model.UpdateUserRequest) (*model.User, error)
	DeleteUser(ctx context.Context, id int64) (bool, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error) {
	return s.userRepo.CreateUser(ctx, req)
}

func (s *userService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}

func (s *userService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}
func (s *userService) UpdateUser(ctx context.Context, id int64, req *model.UpdateUserRequest) (*model.User, error) {
	return s.userRepo.UpdateUser(ctx, id, req)
}

func (s *userService) DeleteUser(ctx context.Context, id int64) (bool, error) {
	return s.userRepo.DeleteUser(ctx, id)
}
