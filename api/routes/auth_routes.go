package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/api/controllers"
	"github.com/yaien/clothes-store-api/api/services"
	"github.com/yaien/clothes-store-api/core"
)

func auth(router *mux.Router, app *core.App) {
	users := services.NewUserService(app.DB)
	tokens := services.NewTokenService(app.Config.JWT, users)
	c := &controllers.AuthController{ Users: users, Tokens: tokens }
	router.HandleFunc("/api/v1/auth/token", c.Token).Methods("POST")
}
