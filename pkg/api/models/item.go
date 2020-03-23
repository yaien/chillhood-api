package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Size struct {
	Label     string `json:"label"`
	Existence int    `json:"existence"`
}

type Item struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `json:"name"`
	Slug        string             `json:"slug"`
	Price       int                `json:"price"`
	Active      bool               `json:"active"`
	Tags        []string           `json:"tags"`
	Pictures    []string           `json:"pictures"`
	Description string             `json:"description"`
	CreatedAt   int64              `json:"createdAt"`
	Sizes       []*Size            `json:"sizes"`
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
