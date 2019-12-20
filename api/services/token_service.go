package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/yaien/clothes-store-api/api/helpers/auth"
	"github.com/yaien/clothes-store-api/core"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type TokenService interface {
	FromPassword(login *auth.Login) (*auth.Response, error)
}

type tokenService struct {
	Users UserService
	Config core.JWTConfig
}


func (s *tokenService) FromPassword(login *auth.Login) (*auth.Response, error) {
	user, err := s.Users.FindOne(bson.M{"email": login.Username})
	if err != nil {
		return nil, err
	}
	if err := user.VerifyPassword(login.Password); err != nil {
		return nil, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Audience:  login.ClientID,
		ExpiresAt: time.Now().Add(s.Config.Duration).Unix(),
		Id:        user.ID.Hex(),
		IssuedAt:  time.Now().Unix(),
		Subject:   "user",
	})

	accessToken, err := token.SignedString(s.Config.Secret)
	if err != nil {
		return nil, err
	}

	response := &auth.Response{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   int(s.Config.Duration.Seconds()),
	}

	return response, nil
}

func NewTokenService(config core.JWTConfig, users UserService) TokenService {
	return &tokenService{users, config}
}