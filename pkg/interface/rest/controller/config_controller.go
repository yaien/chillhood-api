package controller

import (
	response "github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/service"
	"net/http"
)

type ConfigController struct {
	Config service.ConfigService
}

func (cc *ConfigController) Cloudinary(w http.ResponseWriter, _ *http.Request) {
	config := cc.Config.Cloudinary()
	response.JSON(w, map[string]interface{}{
		"cloud":  config.CloudName,
		"preset": config.UploadPreset,
	}, http.StatusOK)
}

func (cc *ConfigController) Epayco(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, cc.Config.Epayco(), http.StatusOK)
}
