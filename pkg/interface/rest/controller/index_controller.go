package controller

import (
	response "github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/response"
	"net/http"
)

// IndexController Controller
type IndexController struct {
}

func (c *IndexController) Get(w http.ResponseWriter, _ *http.Request) {
	response.Send(w, map[string]interface{}{
		"status": "ONLINE",
	})
}
