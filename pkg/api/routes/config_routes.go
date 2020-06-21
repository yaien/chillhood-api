package routes

import (
	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/api/controllers"
)

func config(r *mux.Router, mod *module) {
	config := &controllers.ConfigController{Config: mod.service.config}
	r.HandleFunc("/api/v1/public/config/cloudinary", config.Cloudinary).Methods("GET")
	r.HandleFunc("/api/v1/public/config/epayco", config.Epayco).Methods("GET")

}
