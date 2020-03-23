package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/api/controllers"
)

func index(router *mux.Router, mod *module) {
	ctrl := &controllers.IndexController{}
	router.HandleFunc("/", ctrl.Get)
}
