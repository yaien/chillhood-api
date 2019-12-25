package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/api/controllers"
)

func guest(router *mux.Router, mod *module) {
	controller := &controllers.GuestController{
		Guests: mod.service.guests,
	}

	router.HandleFunc("/api/v1/public/guests", controller.Create).Methods("POST")
	router.Handle("/api/v1/public/guests/{guest_id}", negroni.New(
		negroni.HandlerFunc(controller.Param),
		negroni.WrapFunc(controller.Show),
	)).Methods("GET")

}
