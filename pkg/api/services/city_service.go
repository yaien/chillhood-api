package services

import (
	"context"
	"github.com/yaien/clothes-store-api/pkg/api/models"
)

type CityService interface {
	Search(ctx context.Context, opts models.SearchCityOptions) ([]*models.City, error)
	FindOne(ctx context.Context, opts models.FindOneCityOptions) (*models.City, error)
}

type cityService struct {
	cities models.CityRepository
}

func (c cityService) Search(ctx context.Context, opts models.SearchCityOptions) ([]*models.City, error) {
	return c.cities.Search(ctx, opts)
}

func (c cityService) FindOne(ctx context.Context, opts models.FindOneCityOptions) (*models.City, error) {
	return c.cities.FindOne(ctx, opts)
}

func NewCityService(cities models.CityRepository) CityService {
	return &cityService{cities: cities}
}
