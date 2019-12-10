package controllers

import (
	"net/http"

	"github.com/yaien/clothes-store-api/api/helpers/response"
)

// Index Controller
type Index struct {
}

func (c *Index) Get(w http.ResponseWriter, r *http.Request) {
	response.Send(w, map[string]interface{}{
		"status": "ONLINE",
	})
}
