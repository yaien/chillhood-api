package middlewares

import (
	"net/http"
	"strings"

	"github.com/yaien/clothes-store-api/pkg/api/services"
	"github.com/yaien/clothes-store-api/pkg/core"
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
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := g.Users.FindOneByID(r.Context(), claims.Jti)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ctx := context.WithValue(r.Context(), core.Key("user"), user)
	next(w, r.WithContext(ctx))
}
