package routes

import "github.com/gorilla/mux"

import "github.com/yaien/clothes-store-api/core"

import "github.com/yaien/clothes-store-api/api/controllers"

import "github.com/yaien/clothes-store-api/api/services"

func epayco(router *mux.Router, app *core.App) {
	c := &controllers.EpaycoController{
		Epayco:  services.NewEpaycoService(app.Config.Epayco),
		Invoice: services.NewInvoiceService(app.DB),
		Cart:    services.NewCartService(services.NewItemService(app.DB)),
	}
	router.HandleFunc("/api/v1/epayco/response", c.Response).Methods("GET")
}
