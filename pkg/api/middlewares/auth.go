package middlewares

import (
	"github.com/yaien/clothes-store-api/pkg/api/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"

	"github.com/yaien/clothes-store-api/pkg/api/services"
	"golang.org/x/net/context"
)

// JWTGuard authentication middleware
type JWTGuard struct {
	Tokens services.TokenService
	Users  services.UserService
}

func (g *JWTGuard) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	header := r.Header.Get("Authorization")
	tokenStr := strings.Replace(header, "Bearer ", "", 1)
	claims, err := g.Tokens.Decode(tokenStr)
	if err != nil {
		response.Error(w, &models.Error{Code: "UNAUTHORIZED", Err: err}, http.StatusUnauthorized)
		return
	}

	id, err := primitive.ObjectIDFromHex(claims.Jti)
	if err != nil {
		response.Error(w, &models.Error{Code: "UNAUTHORIZED", Err: err}, http.StatusUnauthorized)
		return
	}

	user, err := g.Users.FindOneByID(r.Context(), id)
	if err != nil {
		response.Error(w, &models.Error{Code: "UNAUTHORIZED", Err: err}, http.StatusUnauthorized)
		return
	}

	ctx := context.WithValue(r.Context(), "user", user)
	next(w, r.WithContext(ctx))
}
