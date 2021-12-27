package controller

import (
	"encoding/json"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/auth"
	response "github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/service"
	"net/http"
)

type AuthController struct {
	Users  service.UserService
	Tokens service.TokenService
}

func (a *AuthController) Token(w http.ResponseWriter, r *http.Request) {
	var login auth.Login
	_ = json.NewDecoder(r.Body).Decode(&login)
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
	user := r.Context().Value("user").(*entity.User)
	response.Send(w, user)
}
