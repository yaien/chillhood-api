package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/api/controllers"
	"github.com/yaien/clothes-store-api/api/services"
	"github.com/yaien/clothes-store-api/core"
)

func cart(router *mux.Router, app *core.App) {
	cart := &controllers.Cart{
		Guests:   services.Guest(app.DB),
		Products: services.Product(app.DB),
	}
	guest := &controllers.Guest{
		Guests: services.Guest(app.DB),
	}

	router.Handle("/api/vi/guests/{guest_id}/cart/products/", negroni.New(
		negroni.HandlerFunc(guest.Param),
		negroni.WrapFunc(cart.AddProduct),
	))

}
