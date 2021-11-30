package models

import (
	"context"
	"errors"
	"time"
)

type Size struct {
	Label     string `json:"label"`
	Existence int    `json:"existence"`
	Body      int    `json:"body"`
	Chest     int    `json:"chest"`
	Sleeve    int    `json:"sleeve"`
}

type Item struct {
	ID          string     `bson:"_id" json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Price       int        `json:"price"`
	Active      bool       `json:"active"`
	Tags        []string   `json:"tags"`
	Pictures    []*Picture `json:"pictures"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	Sizes       []*Size    `json:"sizes"`
}

func (p *Item) Size(label string) (*Size, error) {
	if p.Sizes == nil {
		return nil, errors.New("doesn't have sizes yet")
	}
	for _, size := range p.Sizes {
		if size.Label == label {
			return size, nil
		}
	}
	return nil, errors.New("size not found")
}

type ItemRepository interface {
	Create(ctx context.Context, item *Item) error
	CountByName(ctx context.Context, name string) (int64, error)
	FindOneByID(ctx context.Context, id string) (*Item, error)
	FindOneActiveByID(ctx context.Context, id string) (*Item, error)
	FindOneBySlug(ctx context.Context, slug string) (*Item, error)
	Find(ctx context.Context) ([]*Item, error)
	FindActive(ctx context.Context) ([]*Item, error)
	Update(ctx context.Context, item *Item) error
}
