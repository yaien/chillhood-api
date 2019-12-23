package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/api/controllers"
)

func cart(router *mux.Router, mod *module) {
	cart := &controllers.CartController{
		Guests: mod.service.guests,
		Items:  mod.service.items,
	}
	guest := &controllers.GuestController{
		Guests: mod.service.guests,
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
