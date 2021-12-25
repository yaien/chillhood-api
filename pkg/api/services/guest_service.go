package services

import (
	"context"

	"github.com/yaien/clothes-store-api/pkg/api/models"
)

// GuestService for made crud operations on guest collection
type GuestService interface {
	Create(ctx context.Context, guest *models.Guest) error
	Get(ctx context.Context, id models.ID) (*models.Guest, error)
	Update(ctx context.Context, guest *models.Guest) error
	Reset(ctx context.Context, id models.ID) error
}

type guestService struct {
	repo models.GuestRepository
}

func (s *guestService) Create(ctx context.Context, guest *models.Guest) error {
	err := s.repo.Create(ctx, guest)
	if err != nil {
		return err
	}
	return nil
}

func (s *guestService) Get(ctx context.Context, id models.ID) (*models.Guest, error) {
	return s.repo.FindOneByID(ctx, id)
}

func (s *guestService) Reset(ctx context.Context, id models.ID) error {
	return s.repo.Reset(ctx, id)
}

func (s *guestService) Update(ctx context.Context, guest *models.Guest) error {
	return s.repo.Update(ctx, guest)
}

// NewGuestService return a guest service instance
func NewGuestService(repo models.GuestRepository) GuestService {
	return &guestService{repo: repo}
}
