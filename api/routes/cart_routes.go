package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/api/controllers"
	"github.com/yaien/clothes-store-api/api/services"
	"github.com/yaien/clothes-store-api/core"
)

func cart(router *mux.Router, app *core.App) {
	cart := &controllers.CartController{
		Guests: services.Guest(app.DB),
		Items:  services.NewItemService(app.DB),
	}
	guest := &controllers.GuestController{
		Guests: services.Guest(app.DB),
	}

	router.Handle("/api/v1/guests/{guest_id}/items", negroni.New(
		negroni.HandlerFunc(guest.Param),
		negroni.WrapFunc(cart.Add),
	)).Methods("POST")

	router.Handle("/api/v1/guests/{guest_id}/items/{item_id}", negroni.New(
		negroni.HandlerFunc(guest.Param),
		negroni.WrapFunc(cart.Remove),
	)).Methods("DELETE")

}
