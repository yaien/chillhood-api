package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/controller"
)

func province(router *mux.Router, mod *module) {
	provinces := controller.ProvinceController{
		Provinces: mod.service.provinces,
	}
	router.HandleFunc("/api/v1/public/provinces", provinces.Search).Methods("GET")
}
