package services

import (
	"context"
	"github.com/yaien/clothes-store-api/pkg/api/models"
)

type ProvinceService interface {
	Search(ctx context.Context, opts models.SearchProvinceOptions) ([]*models.Province, error)
}

type provinceService struct {
	provinces models.ProvinceRepository
}

func (p *provinceService) Search(ctx context.Context, opts models.SearchProvinceOptions) ([]*models.Province, error) {
	return p.provinces.Search(ctx, opts)
}

func NewProvinceService(provinces models.ProvinceRepository) ProvinceService {
	return &provinceService{provinces}
}
