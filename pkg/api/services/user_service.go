package services

import (
	"context"
	"time"

	"github.com/yaien/clothes-store-api/pkg/api/models"
)

type UserService interface {
	FindOneByID(ctx context.Context, id string) (*models.User, error)
	FindOneByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
}

type userService struct {
	repo models.UserRepository
}

func (s *userService) FindOneByID(ctx context.Context, id string) (*models.User, error) {
	return s.repo.FindOneByID(ctx, id)
}

func (s *userService) FindOneByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repo.FindOneByEmail(ctx, email)
}

func (s *userService) Create(ctx context.Context, user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return s.repo.Create(ctx, user)
}

func NewUserService(repo models.UserRepository) UserService {
	return &userService{repo}
}
