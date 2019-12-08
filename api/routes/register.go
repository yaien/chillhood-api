package routes

import "github.com/gorilla/mux"

import "github.com/yaien/clothes-store-api/core"

// Register application routes
func Register(app *core.App) *mux.Router {
	router := mux.NewRouter()
	index(router, app)
	guest(router, app)
	cart(router, app)
	product(router, app)

	return router
}
