package models

import "context"

type Province struct {
	ID   string `bson:"_id" json:"id"`
	Name string `json:"name"`
}

type ProvinceRepository interface {
	Search(ctx context.Context, options SearchProvinceOptions) ([]*Province, error)
	FindOneByName(ctx context.Context, name string) (*Province, error)
	Create(ctx context.Context, pr *Province) error
}

type SearchProvinceOptions struct {
	Name  string
	Skip  int64
	Limit int64
}

type City struct {
	ID       string    `bson:"_id" json:"id"`
	Name     string    `json:"name"`
	Shipment int       `json:"shipment"`
	Days     int       `json:"days"`
	Province *Province `json:"province"`
}

type CityRepository interface {
	Search(ctx context.Context, opts SearchCityOptions) ([]*City, error)
	FindOne(ctx context.Context, opts FindOneCityOptions) (*City, error)
	Create(ctx context.Context, city *City) error
	Update(ctx context.Context, city *City) error
}

type FindOneCityOptions struct {
	Name         string
	ProvinceID   string
	ProvinceName string
}
type SearchCityOptions struct {
	Name     string
	Province string
	Skip     int64
	Limit    int64
}
