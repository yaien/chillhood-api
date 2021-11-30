package repository

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCityRepository struct {
	collection *mongo.Collection
}

func (m *MongoCityRepository) FindOne(ctx context.Context, opts models.FindOneCityOptions) (*models.City, error) {
	filter := bson.M{"name": opts.Name}
	if opts.ProvinceID != "" {
		filter["province._id"] = opts.ProvinceID
	}
	var city models.City
	err := m.collection.FindOne(ctx, filter).Decode(&city)
	if err == nil {
		return &city, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &models.Error{Code: "not_found", Err: err}
	}
	return nil, err
}

func (m *MongoCityRepository) Create(ctx context.Context, city *models.City) error {
	city.ID = primitive.NewObjectID().Hex()
	_, err := m.collection.InsertOne(ctx, city)
	return err
}

func (m *MongoCityRepository) Update(ctx context.Context, city *models.City) error {
	_, err := m.collection.UpdateOne(ctx, bson.M{"_id": city.ID}, bson.M{
		"$set": bson.M{
			"name":     city.Name,
			"shipment": city.Shipment,
			"days":     city.Days,
			"province": city.Province,
		},
	})
	return err
}

func NewMongoCityRepository(db *mongo.Database) *MongoCityRepository {
	return &MongoCityRepository{collection: db.Collection("cities")}
}
