package services

import (
	"context"
	"github.com/yaien/clothes-store-api/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	FindOne(filter bson.M) (*models.User, error)
	Create(user *models.User) error
}


type userService struct {
	collection *mongo.Collection
}


func (s *userService) FindOne(filter bson.M) (*models.User, error) {
	var user models.User
	err := s.collection.FindOne(context.TODO(), filter).Decode(&user)
	return &user, err
}

func (s *userService) Create(user *models.User) error {
	user.ID = primitive.NewObjectID()
	_, err := s.collection.InsertOne(context.TODO(), user)
	return err
}


func NewUserService(db *mongo.Database) UserService {
	return &userService{db.Collection("users") }
}
