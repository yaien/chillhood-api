package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/yaien/clothes-store-api/pkg/api/helpers/epayco"
	"github.com/yaien/clothes-store-api/pkg/api/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/api/services"
)

type EpaycoController struct {
	Epayco  services.EpaycoService
	Invoice services.InvoiceService
	Cart    services.CartService
}

func (e *EpaycoController) process(res *epayco.Response) (*models.Invoice, error) {
	if !res.Success {
		return nil, errors.New("UNSUCCESSFULL_RESPONSE")
	}

	if !e.Epayco.Verify(res.Data) {
		return nil, errors.New("INVALID_SIGNATURE")
	}

	invoice, err := e.Invoice.GetByRef(res.Data.Invoice)

	if err != nil {
		return nil, fmt.Errorf("INVOICE_NOT_FOUND: %s", err.Error())
	}

	if invoice.Status != models.Accepted {
		switch res.Data.ResponseCode {
		case epayco.Accepted:
			invoice.Status = models.Accepted
			if !invoice.Cart.Executed {
				if err := e.Cart.Execute(invoice.Cart); err != nil {
					return nil, err
				}
			}
			// create order
			break
		case epayco.Pending:
			invoice.Status = models.Pending
			if !invoice.Cart.Executed {
				if err := e.Cart.Execute(invoice.Cart); err != nil {
					return nil, err
				}
			}
			break
		default:
			invoice.Status = models.Rejected
			if err := e.Cart.Revert(invoice.Cart); err != nil {
				return nil, err
			}
		}
	}

	invoice.Payment = res.Data
	if err := e.Invoice.Update(invoice); err != nil {
		return nil, err
	}

	return invoice, nil
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

	invoice, err := e.process(res)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	response.Send(w, invoice)

}

func (e *EpaycoController) Confirmation(w http.ResponseWriter, r *http.Request) {
	var res epayco.Response
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		response.Error(w, err, http.StatusBadRequest)
	}
	invoice, err := e.process(&res)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	response.Send(w, invoice)
}
