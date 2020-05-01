package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/api/controllers"
)

func config(r *mux.Router, mod *module) {
	controller := &controllers.ConfigController{Config: mod.service.config}
	r.HandleFunc("/v1/config/cloudinary", controller.Cloudinary).Methods("GET")
}
