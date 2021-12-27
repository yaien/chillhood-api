package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/controller"
)

func cart(router *mux.Router, mod *module) {
	cart := &controller.CartController{
		Guests: mod.service.guests,
		Items:  mod.service.items,
	}
	guest := &controller.GuestController{
		Guests: mod.service.guests,
	}

	router.Handle("/api/v1/public/guests/{guest_id}/items", negroni.New(
		negroni.HandlerFunc(guest.Param),
		negroni.WrapFunc(cart.Add),
	)).Methods("POST")

	router.Handle("/api/v1/public/guests/{guest_id}/items/{item_id}", negroni.New(
		negroni.HandlerFunc(guest.Param),
		negroni.WrapFunc(cart.Remove),
	)).Methods("DELETE")

}
