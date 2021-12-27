package service

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/epayco"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/yaien/clothes-store-api/pkg/infrastructure"
)

type EpaycoService interface {
	Request(ref string) (*epayco.Response, error)
	Verify(payment *epayco.Payment) bool
	Process(payment *epayco.Payment) (*entity.Invoice, error)
}

type epaycoService struct {
	invoices InvoiceService
	carts    CartService
	guests   GuestService
	slack    SlackService
	emails   EmailService
	config   *infrastructure.EpaycoConfig
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

func (e *epaycoService) Process(payment *epayco.Payment) (*entity.Invoice, error) {

	if !e.Verify(payment) {
		return nil, errors.New("INVALID_SIGNATURE")
	}

	invoice, err := e.invoices.FindOneByRef(context.TODO(), payment.Invoice)

	if err != nil {
		return nil, fmt.Errorf("INVOICE_NOT_FOUND: %s", err.Error())
	}

	if invoice.Status != entity.Accepted {
		switch payment.ResponseCode {
		case epayco.Accepted:
			invoice.Status = entity.Accepted
			invoice.Shipping.Status = entity.Preparing
			if !invoice.Cart.Executed {
				if err := e.carts.Execute(invoice.Cart); err != nil {
					return nil, err
				}
			}
			if err := e.guests.Reset(context.TODO(), invoice.GuestID); err != nil {
				return nil, err
			}
			e.slack.NotifySale(invoice)
			e.emails.NotifySale(invoice)
		case epayco.Pending:
			invoice.Status = entity.Pending
			if !invoice.Cart.Executed {
				if err := e.carts.Execute(invoice.Cart); err != nil {
					return nil, err
				}
			}
			if err := e.guests.Reset(context.TODO(), invoice.GuestID); err != nil {
				return nil, err
			}
		default:
			invoice.Status = entity.Rejected
			if err := e.carts.Revert(invoice.Cart); err != nil {
				return nil, err
			}
		}
	}

	invoice.Payment = payment
	if err := e.invoices.Update(context.TODO(), invoice); err != nil {
		return nil, err
	}

	return invoice, nil
}

func NewEpaycoService(
	config *infrastructure.EpaycoConfig,
	baseURL *url.URL,
	invoiceSrv InvoiceService,
	cartSrv CartService,
	guestSrv GuestService,
	slackSrv SlackService,
	emailSrv EmailService) EpaycoService {
	return &epaycoService{
		invoiceSrv,
		cartSrv,
		guestSrv,
		slackSrv,
		emailSrv,
		config,
		baseURL,
	}
}
