package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/pkg/api/controllers"
)

func auth(router *mux.Router, mod *module) {
	controller := &controllers.AuthController{
		Users:  mod.service.users,
		Tokens: mod.service.tokens,
	}
	router.HandleFunc("/api/v1/auth/token", controller.Token).Methods("POST")
	router.Handle("/api/v1/user", negroni.New(
		mod.middleware.jwt,
		negroni.WrapFunc(controller.User),
	))
}
