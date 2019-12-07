package routes

import "github.com/gorilla/mux"

import "github.com/yaien/clothes-store-api/core"

func Register(app *core.App) *mux.Router {
	router := mux.NewRouter()
	index(router, app)
	return router
}
