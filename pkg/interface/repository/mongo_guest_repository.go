package repository

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoGuestRepository struct {
	collection *mongo.Collection
}

func (m *MongoGuestRepository) Reset(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$unset": bson.M{"cart": ""}}
	_, err := m.collection.UpdateOne(ctx, filter, update)
	return err
}

func (m *MongoGuestRepository) Create(ctx context.Context, guest *models.Guest) error {
	guest.ID = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, guest)
	return err
}

func (m *MongoGuestRepository) FindOneByID(ctx context.Context, id primitive.ObjectID) (*models.Guest, error) {
	var guest models.Guest
	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&guest)
	if err == nil {
		return &guest, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &models.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *MongoGuestRepository) Update(ctx context.Context, guest *models.Guest) error {
	_, err := m.collection.UpdateOne(ctx, bson.M{"_id": guest.ID}, bson.M{
		"$set": bson.M{"cart": guest.Cart},
	})
	return err
}

// NewMongoGuestRepository returns a new guest repository using mongodb
func NewMongoGuestRepository(db *mongo.Database) *MongoGuestRepository {
	return &MongoGuestRepository{collection: db.Collection("guests")}
}
