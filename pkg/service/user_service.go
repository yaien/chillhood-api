package service

import (
	"context"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"time"
)

type UserService interface {
	FindOneByID(ctx context.Context, id entity.ID) (*entity.User, error)
	FindOneByEmail(ctx context.Context, email string) (*entity.User, error)
	FindReportable(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
}

type userService struct {
	repo entity.UserRepository
}

func (s *userService) FindOneByID(ctx context.Context, id entity.ID) (*entity.User, error) {
	return s.repo.FindOneByID(ctx, id)
}

func (s *userService) FindOneByEmail(ctx context.Context, email string) (*entity.User, error) {
	return s.repo.FindOneByEmail(ctx, email)
}

func (s *userService) FindReportable(ctx context.Context) ([]*entity.User, error) {
	return s.repo.FindReportable(ctx)
}

func (s *userService) Create(ctx context.Context, user *entity.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return s.repo.Create(ctx, user)
}

func NewUserService(repo entity.UserRepository) UserService {
	return &userService{repo}
}
