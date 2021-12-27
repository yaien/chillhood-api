package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/controller"
)

func config(r *mux.Router, mod *module) {
	config := &controller.ConfigController{Config: mod.service.config}
	r.HandleFunc("/api/v1/public/config/cloudinary", config.Cloudinary).Methods("GET")
	r.HandleFunc("/api/v1/public/config/epayco", config.Epayco).Methods("GET")

}
