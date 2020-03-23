package controllers

import (
	"net/http"

	"github.com/yaien/clothes-store-api/pkg/api/helpers/response"
)

// IndexController Controller
type IndexController struct {
}

func (c *IndexController) Get(w http.ResponseWriter, r *http.Request) {
	response.Send(w, map[string]interface{}{
		"status": "ONLINE",
	})
}
