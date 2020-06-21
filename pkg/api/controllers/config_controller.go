package controllers

import (
	"net/http"

	"github.com/yaien/clothes-store-api/pkg/api/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/api/services"
)

type ConfigController struct {
	Config services.ConfigService
}

func (cc *ConfigController) Cloudinary(w http.ResponseWriter, r *http.Request) {
	config := cc.Config.Cloudinary()
	response.JSON(w, map[string]interface{}{
		"cloud":  config.CloudName,
		"preset": config.UploadPreset,
	}, http.StatusOK)
}

func (cc *ConfigController) Epayco(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, cc.Config.Epayco(), http.StatusOK)
}
