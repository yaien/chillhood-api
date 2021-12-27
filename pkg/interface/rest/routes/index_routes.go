package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/controller"
)

func index(router *mux.Router, mod *module) {
	ctrl := &controller.IndexController{}
	router.HandleFunc("/", ctrl.Get)
}
