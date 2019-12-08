package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string
	Price       int
	Active      bool
	Tags        []string
	Pictures    []string
	Description string
	CreatedAt   int
	Sizes       []Size
}

type Size struct {
	Label     string
	Existence int
}
