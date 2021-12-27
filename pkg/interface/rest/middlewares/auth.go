package middlewares

import (
	"github.com/yaien/clothes-store-api/pkg/entity"
	response2 "github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"

	"golang.org/x/net/context"
)

// JWTGuard authentication middleware
type JWTGuard struct {
	Tokens service.TokenService
	Users  service.UserService
}

func (g *JWTGuard) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	header := r.Header.Get("Authorization")
	tokenStr := strings.Replace(header, "Bearer ", "", 1)
	claims, err := g.Tokens.Decode(tokenStr)
	if err != nil {
		response2.Error(w, &entity.Error{Code: "UNAUTHORIZED", Err: err}, http.StatusUnauthorized)
		return
	}

	id, err := primitive.ObjectIDFromHex(claims.Jti)
	if err != nil {
		response2.Error(w, &entity.Error{Code: "UNAUTHORIZED", Err: err}, http.StatusUnauthorized)
		return
	}

	user, err := g.Users.FindOneByID(r.Context(), id)
	if err != nil {
		response2.Error(w, &entity.Error{Code: "UNAUTHORIZED", Err: err}, http.StatusUnauthorized)
		return
	}

	ctx := context.WithValue(r.Context(), "user", user)
	next(w, r.WithContext(ctx))
}
