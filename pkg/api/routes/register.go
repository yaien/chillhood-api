package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/core"
)

// Register application routes
func Register(app *core.App) *mux.Router {
	router := mux.NewRouter()
	mod := bundle(app)
	index(router, mod)
	city(router, mod)
	province(router, mod)
	guest(router, mod)
	auth(router, mod)
	cart(router, mod)
	item(router, mod)
	invoice(router, mod)
	epayco(router, mod)
	config(router, mod)
	return router
}
