package service

import (
	"context"
	"github.com/yaien/clothes-store-api/pkg/entity"
)

type ProvinceService interface {
	Search(ctx context.Context, opts entity.SearchProvinceOptions) ([]*entity.Province, error)
}

type provinceService struct {
	provinces entity.ProvinceRepository
}

func (p *provinceService) Search(ctx context.Context, opts entity.SearchProvinceOptions) ([]*entity.Province, error) {
	return p.provinces.Search(ctx, opts)
}

func NewProvinceService(provinces entity.ProvinceRepository) ProvinceService {
	return &provinceService{provinces}
}
