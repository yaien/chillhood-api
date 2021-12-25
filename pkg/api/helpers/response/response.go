package response

import (
	"encoding/json"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"net/http"
)

// JSON -> Send a json response
func JSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Send -> Send a status ok json response
func Send(w http.ResponseWriter, data interface{}) {
	JSON(w, data, http.StatusOK)
}

// Error -> Send an error json response
func Error(w http.ResponseWriter, err error, status int) {
	var e *models.Error
	if errors.As(err, &e) {
		JSON(w, map[string]string{"error": e.Code}, status)
		return
	}

	JSON(w, map[string]string{"error": err.Error()}, status)
}
