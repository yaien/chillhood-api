package models

import "context"

type Guest struct {
	ID   string `bson:"_id" json:"id"`
	Cart *Cart  `json:"cart,omitempty"`
}

type GuestRepository interface {
	Create(ctx context.Context, guest *Guest) error
	FindOneByID(ctx context.Context, id string) (*Guest, error)
	Update(ctx context.Context, guest *Guest) error
	Reset(ctx context.Context, id string) error
}
