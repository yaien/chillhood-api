package repository

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func (m *MongoUserRepository) FindOneByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == nil {
		return &user, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &models.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *MongoUserRepository) FindOneByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := m.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == nil {
		return &user, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &models.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *MongoUserRepository) Create(ctx context.Context, u *models.User) error {
	u.ID = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, u)
	return err
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{collection: db.Collection("users")}
}
