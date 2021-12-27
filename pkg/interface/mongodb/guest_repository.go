package mongodb

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GuestRepository struct {
	collection *mongo.Collection
}

func (m *GuestRepository) Reset(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$unset": bson.M{"cart": ""}}
	_, err := m.collection.UpdateOne(ctx, filter, update)
	return err
}

func (m *GuestRepository) Create(ctx context.Context, guest *entity.Guest) error {
	guest.ID = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, guest)
	return err
}

func (m *GuestRepository) FindOneByID(ctx context.Context, id primitive.ObjectID) (*entity.Guest, error) {
	var guest entity.Guest
	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&guest)
	if err == nil {
		return &guest, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &entity.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *GuestRepository) Update(ctx context.Context, guest *entity.Guest) error {
	_, err := m.collection.UpdateOne(ctx, bson.M{"_id": guest.ID}, bson.M{
		"$set": bson.M{"cart": guest.Cart},
	})
	return err
}

// NewGuestRepository returns a new guest repository using mongodb
func NewGuestRepository(db *mongo.Database) *GuestRepository {
	return &GuestRepository{collection: db.Collection("guests")}
}
