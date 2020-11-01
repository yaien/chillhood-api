package controllers

import (
	"errors"
	"fmt"
	"github.com/yaien/clothes-store-api/pkg/api/helpers/epayco"
	"log"
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

	if !res.Success {
		response.Error(w, errors.New("UNSUCCESSFULL_RESPONSE"), http.StatusBadRequest)
		return
	}

	invoice, err := e.Epayco.Process(res.Data)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	response.Send(w, invoice)

}

func (e *EpaycoController) Confirmation(w http.ResponseWriter, r *http.Request) {
	payment := epayco.ParsePaymentFromRequest(r)
	invoice, err := e.Epayco.Process(payment)
	if err != nil {
		e := fmt.Errorf("failed proccessing: %w", err)
		log.Println("epayco confirmation:", e.Error())
		response.Error(w, e, http.StatusBadRequest)
		return
	}
	response.Send(w, invoice)
}
