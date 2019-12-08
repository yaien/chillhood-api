package core

import (
	"encoding/json"
	"net/http"
)

// Controller -> Base Controller Struct
type Controller struct {
	App *App
}

// JSON -> Send a json response
func (c *Controller) JSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Send -> Send a status ok json response
func (c *Controller) Send(w http.ResponseWriter, data interface{}) {
	c.JSON(w, data, http.StatusOK)
}

func (c *Controller) Error(w http.ResponseWriter, err error, status int) {
	c.JSON(w, map[string]string{"error": err.Error()}, status)
}
