package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/api/controllers"
	"github.com/yaien/clothes-store-api/core"
)

func index(router *mux.Router, app *core.App) {
	ctrl := &controllers.Index{Controller: core.Controller{App: app}}
	router.HandleFunc("/", ctrl.Get)
}
