package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/yaien/clothes-store-api/pkg/api/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/api/services"
)

type EpaycoController struct {
	Epayco services.EpaycoService
}

func (e *EpaycoController) Response(w http.ResponseWriter, r *http.Request) {
	ref := r.URL.Query().Get("ref_payco")
	if len(ref) == 0 {
		response.Error(w, errors.New("MISSING_REF"), http.StatusBadRequest)
		return
	}

	res, err := e.Epayco.Request(ref)

	if err != nil {
		response.Error(w, fmt.Errorf("REF_NOT_FOUND: %s", err.Error()), http.StatusNotFound)
		return
	}

	invoice, err := e.Epayco.Process(res)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	response.Send(w, invoice)

}

func (e *EpaycoController) Confirmation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		response.Error(w, fmt.Errorf("failed parsing form: %w", err), http.StatusBadRequest)
		return
	}
	ref := r.Form.Get("x_ref_payco")
	res, err := e.Epayco.Request(ref)
	if err != nil {
		response.Error(w, fmt.Errorf("REF_NOT_FOUND: %s", err.Error()), http.StatusNotFound)
		return
	}
	invoice, err := e.Epayco.Process(res)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}
	response.Send(w, invoice)
}
