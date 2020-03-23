package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/yaien/clothes-store-api/pkg/api/helpers/auth"
	"github.com/yaien/clothes-store-api/pkg/api/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/api/services"
	"github.com/yaien/clothes-store-api/pkg/core"
)

type AuthController struct {
	Users  services.UserService
	Tokens services.TokenService
}

func (a *AuthController) Token(w http.ResponseWriter, r *http.Request) {
	var login auth.Login
	json.NewDecoder(r.Body).Decode(&login)
	switch login.GrantType {
	case "password":
		res, err := a.Tokens.FromPassword(&login)
		if err != nil {
			response.Error(w, err, http.StatusUnauthorized)
			return
		}
		response.Send(w, res)
	default:
		response.Error(w, errors.New("INVALID_GRANT_TYPE"), http.StatusBadRequest)
	}
}

func (a *AuthController) User(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(core.Key("user")).(*models.User)
	response.Send(w, user)
}
