package controllers

import (
	"errors"
	"log"
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

func (e *EpaycoController) Response(w http.ResponseWriter, r *http.Request) {
	ref := r.URL.Query().Get("ref_payco")
	if len(ref) == 0 {
		response.Error(w, errors.New("MISSING_REF"), http.StatusBadRequest)
		return
	}

	res, err := e.Epayco.Request(ref)

	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	if !res.Success {
		response.Error(w, errors.New("UNSUCCESSFULL_RESPONSE"), http.StatusBadRequest)
		return
	}

	if !e.Epayco.Verify(res.Data) {
		response.Error(w, errors.New("INVALID_SIGNATURE"), http.StatusBadRequest)
		return
	}

	invoice, err := e.Invoice.GetByRef(res.Data.Invoice)

	if err != nil {
		response.Error(w, errors.New("INVOICE_NOT_FOUND"), http.StatusBadRequest)
		return
	}

	if invoice.Status != models.Accepted {
		switch res.Data.ResponseCode {
		case epayco.Accepted:
			invoice.Status = models.Accepted
			if !invoice.Cart.Executed {
				if err := e.Cart.Execute(invoice.Cart); err != nil {
					response.Error(w, err, http.StatusBadRequest)
					return
				}
			}
			// create order
			break
		case epayco.Pending:
			invoice.Status = models.Pending
			if !invoice.Cart.Executed {
				if err := e.Cart.Execute(invoice.Cart); err != nil {
					response.Error(w, err, http.StatusBadRequest)
					return
				}
			}
			break
		default:
			invoice.Status = models.Rejected
			if err := e.Cart.Revert(invoice.Cart); err != nil {
				log.Println(err)
				response.Error(w, err, http.StatusBadRequest)
				return
			}
		}
	}

	invoice.Payment = res.Data
	if err := e.Invoice.Update(invoice); err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	response.Send(w, invoice)

}

func (e *EpaycoController) Checkout(w http.ResponseWriter, r *http.Request) {
	response.Send(w, e.Epayco.CheckoutArgs())
}
