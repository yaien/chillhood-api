package controllers

import "github.com/yaien/clothes-store-api/core"

import "net/http"

// Index Controller
type Index struct {
	core.Controller
}

func (c *Index) Get(w http.ResponseWriter, r *http.Request) {
	c.Send(w, map[string]interface{}{
		"status": "ONLINE",
	})
}
