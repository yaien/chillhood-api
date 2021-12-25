package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/o1egl/paseto"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"time"

	"github.com/yaien/clothes-store-api/pkg/api/helpers/auth"
	"github.com/yaien/clothes-store-api/pkg/core"
)

type TokenService interface {
	FromPassword(login *auth.Login) (*auth.Response, error)
	Decode(token string) (*paseto.JSONToken, error)
}

type tokenService struct {
	Users  UserService
	Config *core.JWTConfig
	Client *core.ClientConfig
}

func (s *tokenService) isClientKeyValid(clientKey string) bool {
	for _, key := range s.Client.Keys {
		if key == clientKey {
			return true
		}
	}
	return false
}

func (s *tokenService) FromPassword(login *auth.Login) (*auth.Response, error) {

	if !s.isClientKeyValid(login.ClientID) {
		return nil, errors.New("INVALID_CLIENT_CREDENTIALS")
	}

	user, err := s.Users.FindOneByEmail(context.TODO(), login.Username)
	if err != nil {
		return nil, fmt.Errorf("failed finding user: %w", err)
	}
	if err := user.VerifyPassword(login.Password); err != nil {
		return nil, &models.Error{Code: "INVALID_PASSWORD", Err: err}
	}

	claims := paseto.JSONToken{
		Audience:   login.ClientID,
		Expiration: time.Now().Add(s.Config.Duration),
		IssuedAt:   time.Now(),
		Jti:        user.ID.Hex(),
	}

	v2 := paseto.V2{}
	token, err := v2.Encrypt(s.Config.Secret, claims, nil)
	if err != nil {
		return nil, fmt.Errorf("failed encrypt: %w", err)
	}

	response := &auth.Response{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int(s.Config.Duration.Seconds()),
	}

	return response, nil
}

func (s *tokenService) Decode(token string) (*paseto.JSONToken, error) {
	var claims paseto.JSONToken
	var v2 paseto.V2
	err := v2.Decrypt(token, s.Config.Secret, &claims, nil)
	if err != nil {
		return nil, err
	}
	return &claims, nil
}

func NewTokenService(client *core.ClientConfig, config *core.JWTConfig, users UserService) TokenService {
	return &tokenService{users, config, client}
}
