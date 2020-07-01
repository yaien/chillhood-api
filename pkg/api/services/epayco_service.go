package services

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/yaien/clothes-store-api/pkg/api/helpers/epayco"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/core"
)

type EpaycoService interface {
	Request(ref string) (*epayco.Response, error)
	Verify(payment *epayco.Payment) bool
	Process(res *epayco.Response) (*models.Invoice, error)
}

type epaycoService struct {
	invoices InvoiceService
	carts    CartService
	guests   GuestService
	config   *core.EpaycoConfig
	baseURL  *url.URL
}

func (e *epaycoService) Request(ref string) (*epayco.Response, error) {
	res, err := http.Get("https://secure.epayco.co/validation/v1/reference/" + ref)
	if err != nil {
		return nil, err
	}
	var response epayco.Response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (e *epaycoService) Verify(payment *epayco.Payment) bool {
	payload := []string{
		e.config.CustomerID,
		e.config.Key,
		strconv.Itoa(payment.Ref),
		payment.TransactionID,
		strconv.Itoa(payment.Amount),
		payment.CurrencyCode,
	}
	source := strings.Join(payload, "^")
	signature := fmt.Sprintf("%x", sha256.Sum256([]byte(source)))
	return signature == payment.Signature
}

func (e *epaycoService) Process(res *epayco.Response) (*models.Invoice, error) {
	if !res.Success {
		return nil, errors.New("UNSUCCESSFULL_RESPONSE")
	}

	if !e.Verify(res.Data) {
		return nil, errors.New("INVALID_SIGNATURE")
	}

	invoice, err := e.invoices.GetByRef(res.Data.Invoice)

	if err != nil {
		return nil, fmt.Errorf("INVOICE_NOT_FOUND: %s", err.Error())
	}

	if invoice.Status != models.Accepted {
		switch res.Data.ResponseCode {
		case epayco.Accepted:
			invoice.Status = models.Accepted
			invoice.Shipping.Status = models.Preparing
			if !invoice.Cart.Executed {
				if err := e.carts.Execute(invoice.Cart); err != nil {
					return nil, err
				}
			}
			if err := e.guests.Reset(invoice.GuestID.Hex()); err != nil {
				return nil, err
			}
			break
		case epayco.Pending:
			invoice.Status = models.Pending
			if !invoice.Cart.Executed {
				if err := e.carts.Execute(invoice.Cart); err != nil {
					return nil, err
				}
			}
			if err := e.guests.Reset(invoice.GuestID.Hex()); err != nil {
				return nil, err
			}
			break
		default:
			invoice.Status = models.Rejected
			if err := e.carts.Revert(invoice.Cart); err != nil {
				return nil, err
			}
		}
	}

	invoice.Payment = res.Data
	if err := e.invoices.Update(invoice); err != nil {
		return nil, err
	}

	return invoice, nil
}

func NewEpaycoService(
	config *core.EpaycoConfig,
	baseURL *url.URL,
	invoiceSrv InvoiceService,
	cartSrv CartService,
	guestSrv GuestService) EpaycoService {
	return &epaycoService{invoiceSrv, cartSrv, guestSrv, config, baseURL}
}
