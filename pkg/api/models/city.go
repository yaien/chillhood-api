package models

import "context"

type Province struct {
	ID   string `bson:"_id" json:"id"`
	Name string `json:"name"`
}

type ProvinceRepository interface {
	FindOneByName(ctx context.Context, name string) (*Province, error)
	Create(ctx context.Context, pr *Province) error
}

type City struct {
	ID       string    `bson:"_id" json:"id"`
	Name     string    `json:"name"`
	Shipment int       `json:"shipment"`
	Days     int       `json:"days"`
	Province *Province `json:"province"`
}

type CityRepository interface {
	FindOne(ctx context.Context, opts FindOneCityOptions) (*City, error)
	Create(ctx context.Context, city *City) error
	Update(ctx context.Context, city *City) error
}

type FindOneCityOptions struct {
	Name       string
	ProvinceID string
}
