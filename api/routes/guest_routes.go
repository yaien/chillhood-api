package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/api/controllers"
	"github.com/yaien/clothes-store-api/api/services"
	"github.com/yaien/clothes-store-api/core"
)

func guest(router *mux.Router, app *core.App) {
	controller := &controllers.Guest{
		Guests: services.Guest(app.DB),
	}

	router.HandleFunc("/api/v1/guests", controller.Create).Methods("POST")
	router.Handle("/api/v1/guests/{guest_id}", negroni.New(
		negroni.HandlerFunc(controller.Param),
		negroni.WrapFunc(controller.Show),
	)).Methods("GET")

}
