package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/core"
)

// Register application routes
func Register(app *core.App) *mux.Router {
	router := mux.NewRouter()
	index(router, app)
	guest(router, app)
	cart(router, app)
	item(router, app)
	invoice(router, app)

	return router
}
