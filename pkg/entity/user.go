package entity

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ID = primitive.ObjectID

type User struct {
	ID        ID        `bson:"_id" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) HashPassword() error {
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(password)
	return nil
}

func (u *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

type UserRepository interface {
	FindOneByID(ctx context.Context, id ID) (*User, error)
	FindOneByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, u *User) error
}
