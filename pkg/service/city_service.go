package service

import (
	"context"
	"github.com/yaien/clothes-store-api/pkg/entity"
)

type CityService interface {
	Search(ctx context.Context, opts entity.SearchCityOptions) ([]*entity.City, error)
	FindOne(ctx context.Context, opts entity.FindOneCityOptions) (*entity.City, error)
}

type cityService struct {
	cities entity.CityRepository
}

func (c cityService) Search(ctx context.Context, opts entity.SearchCityOptions) ([]*entity.City, error) {
	return c.cities.Search(ctx, opts)
}

func (c cityService) FindOne(ctx context.Context, opts entity.FindOneCityOptions) (*entity.City, error) {
	return c.cities.FindOne(ctx, opts)
}

func NewCityService(cities entity.CityRepository) CityService {
	return &cityService{cities: cities}
}
