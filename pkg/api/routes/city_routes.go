package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/api/controllers"
)

func city(router *mux.Router, mod *module) {
	cities := controllers.CityController{
		Cities: mod.service.cities,
	}
	router.HandleFunc("/api/v1/public/cities", cities.Search).Methods("GET")
}
