package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/api/controllers"
)

func epayco(router *mux.Router, mod *module) {
	c := &controllers.EpaycoController{
		Epayco:  mod.service.epayco,
		Invoice: mod.service.invoices,
		Cart:    mod.service.carts,
	}
	router.HandleFunc("/api/v1/epayco/response", c.Response).Methods("GET")
	router.HandleFunc("/api/v1/epayco/checkout", c.Checkout).Methods("GET")
}
