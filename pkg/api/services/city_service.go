package services

import (
	"context"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SearchCityOptions struct {
	Name     string
	Province string
	Skip     int64
	Limit    int64
}

type CityService interface {
	Search(opts SearchCityOptions) ([]*models.City, error)
	FindOneByID(id primitive.ObjectID) (*models.City, error)
}

type cityService struct {
	collection *mongo.Collection
}

func (c cityService) Search(opts SearchCityOptions) ([]*models.City, error) {
	var cities []*models.City
	filter := bson.M{
		"name":          primitive.Regex{Pattern: opts.Name, Options: "ig"},
		"province.name": primitive.Regex{Pattern: opts.Province, Options: "ig"},
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
	cursor, err := c.collection.Find(context.TODO(), filter, ops)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var city models.City
		err = cursor.Decode(&city)
		if err != nil {
			return nil, err
		}
		cities = append(cities, &city)
	}
	return cities, nil
}

func (c cityService) FindOneByID(id primitive.ObjectID) (*models.City, error) {
	var city models.City
	err := c.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&city)
	if err != nil {
		return nil, err
	}
	return &city, nil
}

func NewCityService(db *mongo.Database) CityService {
	return &cityService{collection: db.Collection("cities")}
}
