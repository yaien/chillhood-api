package validation

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/api/helpers/response"
)

type Input interface {
	Validate() map[string]string
}

type Key string

func NewHandler(input Input) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		data := reflect.ValueOf(input).Interface().(Input)
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			response.Error(w, err, http.StatusBadRequest)
			return
		}
		errors := data.Validate()
		if len(errors) > 0 {
			response.JSON(w, errors, http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), Key("data"), data)
		next(w, r.WithContext(ctx))
	}
}
