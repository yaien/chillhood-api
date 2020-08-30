package services

import (
	"context"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SearchProvinceOptions struct {
	Name  string
	Skip  int64
	Limit int64
}

type ProvinceService interface {
	Search(opts SearchProvinceOptions) ([]*models.Province, error)
}

type provinceService struct {
	collection *mongo.Collection
}

func (p provinceService) Search(opts SearchProvinceOptions) ([]*models.Province, error) {
	var provinces []*models.Province
	filter := bson.M{
		"name": primitive.Regex{Pattern: opts.Name, Options: "i"},
	}
	skip := opts.Skip
	if skip < 0 {
		skip = 0
	}
	limit := opts.Limit
	if limit < 1 || limit > 100 {
		limit = 10
	}
	ops := options.Find().SetSkip(skip).SetLimit(limit)
	cursor, err := p.collection.Find(context.TODO(), filter, ops)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var province models.Province
		err = cursor.Decode(&province)
		if err != nil {
			return nil, err
		}
		provinces = append(provinces, &province)
	}
	return provinces, nil
}

func NewProvinceService(db *mongo.Database) ProvinceService {
	return &provinceService{db.Collection("provinces")}
}
