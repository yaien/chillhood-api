package services

import (
	"context"

	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GuestService for made crud operations on guest collection
type GuestService interface {
	Create(guest *models.Guest) error
	Get(id string) (*models.Guest, error)
	Update(guest *models.Guest) error
	Reset(id string) error
}

type guestService struct {
	collection *mongo.Collection
}

func (s *guestService) Create(guest *models.Guest) error {
	guest.ID = primitive.NewObjectID()
	_, err := s.collection.InsertOne(context.TODO(), guest)
	if err != nil {
		return err
	}
	return nil
}

func (s *guestService) Get(id string) (*models.Guest, error) {
	var guest models.Guest
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": _id}
	err = s.collection.FindOne(context.TODO(), filter).Decode(&guest)
	if err != nil {
		return nil, err
	}
	return &guest, nil
}

func (s *guestService) Reset(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"cart": nil}
	_, err = s.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (s *guestService) Update(guest *models.Guest) error {
	filter := bson.M{"_id": guest.ID}
	update := bson.M{"$set": guest}
	_, err := s.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return err
}

// NewGuestService return a guest service instance
func NewGuestService(db *mongo.Database) GuestService {
	return &guestService{
		collection: db.Collection("guests"),
	}
}
