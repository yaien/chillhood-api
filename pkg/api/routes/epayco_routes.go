package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/api/controllers"
)

func epayco(router *mux.Router, mod *module) {
	c := &controllers.EpaycoController{
		Epayco: mod.service.epayco,
	}
	router.HandleFunc("/api/v1/public/epayco/response", c.Response).Methods("GET", "POST")
	router.HandleFunc("/api/v1/public/epayco/confirmation", c.Confirmation).Methods("POST")
}
