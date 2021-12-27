package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/controller"
)

func city(router *mux.Router, mod *module) {
	cities := controller.CityController{
		Cities: mod.service.cities,
	}
	router.HandleFunc("/api/v1/public/cities", cities.Search).Methods("GET")
}
