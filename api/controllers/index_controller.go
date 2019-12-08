package controllers

import "net/http"

// Index Controller
type Index struct {
	*controller
}

func (c *Index) Get(w http.ResponseWriter, r *http.Request) {
	c.Send(w, map[string]interface{}{
		"status": "ONLINE",
	})
}
