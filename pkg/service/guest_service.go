package service

import (
	"context"
	"github.com/yaien/clothes-store-api/pkg/entity"
)

// GuestService for made crud operations on guest collection
type GuestService interface {
	Create(ctx context.Context, guest *entity.Guest) error
	Get(ctx context.Context, id entity.ID) (*entity.Guest, error)
	Update(ctx context.Context, guest *entity.Guest) error
	Reset(ctx context.Context, id entity.ID) error
}

type guestService struct {
	repo entity.GuestRepository
}

func (s *guestService) Create(ctx context.Context, guest *entity.Guest) error {
	err := s.repo.Create(ctx, guest)
	if err != nil {
		return err
	}
	return nil
}

func (s *guestService) Get(ctx context.Context, id entity.ID) (*entity.Guest, error) {
	return s.repo.FindOneByID(ctx, id)
}

func (s *guestService) Reset(ctx context.Context, id entity.ID) error {
	return s.repo.Reset(ctx, id)
}

func (s *guestService) Update(ctx context.Context, guest *entity.Guest) error {
	return s.repo.Update(ctx, guest)
}

// NewGuestService return a guest service instance
func NewGuestService(repo entity.GuestRepository) GuestService {
	return &guestService{repo: repo}
}
