package mongodb

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func (m *UserRepository) FindOneByID(ctx context.Context, id primitive.ObjectID) (*entity.User, error) {
	var user entity.User
	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == nil {
		return &user, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &entity.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *UserRepository) FindOneByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := m.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == nil {
		return &user, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &entity.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *UserRepository) Create(ctx context.Context, u *entity.User) error {
	u.ID = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, u)
	return err
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{collection: db.Collection("users")}
}
