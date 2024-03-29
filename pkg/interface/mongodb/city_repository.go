package mongodb

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CityRepository struct {
	collection *mongo.Collection
}

func (m *CityRepository) Search(ctx context.Context, opts entity.SearchCityOptions) ([]*entity.City, error) {
	var cities []*entity.City
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
	cursor, err := m.collection.Find(ctx, filter, ops)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var city entity.City
		err = cursor.Decode(&city)
		if err != nil {
			return nil, err
		}
		cities = append(cities, &city)
	}
	if cities == nil {
		cities = make([]*entity.City, 0)
	}
	return cities, nil
}

func (m *CityRepository) FindOne(ctx context.Context, opts entity.FindOneCityOptions) (*entity.City, error) {
	filter := bson.M{"name": opts.Name}
	if !opts.ProvinceID.IsZero() {
		filter["province._id"] = opts.ProvinceID
	}
	if opts.ProvinceName != "" {
		filter["province.name"] = opts.ProvinceName
	}
	var city entity.City
	err := m.collection.FindOne(ctx, filter).Decode(&city)
	if err == nil {
		return &city, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &entity.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *CityRepository) Create(ctx context.Context, city *entity.City) error {
	city.ID = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, city)
	return err
}

func (m *CityRepository) Update(ctx context.Context, city *entity.City) error {
	_, err := m.collection.UpdateOne(ctx, bson.M{"_id": city.ID}, bson.M{"$set": city})
	return err
}

func NewCityRepository(db *mongo.Database) *CityRepository {
	return &CityRepository{collection: db.Collection("cities")}
}
