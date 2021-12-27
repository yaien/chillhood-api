package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/controller"
)

func guest(router *mux.Router, mod *module) {
	ctrl := &controller.GuestController{
		Guests: mod.service.guests,
	}

	router.HandleFunc("/api/v1/public/guests", ctrl.Create).Methods("POST")
	router.Handle("/api/v1/public/guests/{guest_id}", negroni.New(
		negroni.HandlerFunc(ctrl.Param),
		negroni.WrapFunc(ctrl.Show),
	)).Methods("GET")

}
