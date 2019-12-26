package middlewares

import (
	"net/http"
	"strings"

	"github.com/yaien/clothes-store-api/api/services"
	"github.com/yaien/clothes-store-api/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type JWTGuard struct {
	Tokens services.TokenService
	Users  services.UserService
}

func (g *JWTGuard) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	header := r.Header.Get("Authorization")
	tokenStr := strings.Replace(header, "Bearer ", "", 1)
	claims, err := g.Tokens.Decode(tokenStr)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id, err := primitive.ObjectIDFromHex(claims.Id)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user, err := g.Users.FindOne(bson.M{"_id": id})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ctx := context.WithValue(r.Context(), core.Key("user"), user)
	next(w, r.WithContext(ctx))
}
