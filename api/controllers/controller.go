package controllers

import (
	"encoding/json"
	"net/http"
)

type key string

// Controller -> Base Controller Struct
type controller struct {
}

// JSON -> Send a json response
func (c *controller) JSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}

// Send -> Send a status ok json response
func (c *controller) Send(w http.ResponseWriter, data interface{}) {
	c.JSON(w, data, http.StatusOK)
}

func (c *controller) Error(w http.ResponseWriter, err error, status int) {
	c.JSON(w, map[string]string{"error": err.Error()}, status)
}
