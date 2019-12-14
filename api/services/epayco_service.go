package services

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/yaien/clothes-store-api/api/helpers/epayco"
	"github.com/yaien/clothes-store-api/core"
)

type EpaycoService interface {
	Request(ref string) (*epayco.Response, error)
	Verify(payment *epayco.Payment) bool
}

type epaycoService struct {
	config core.EpaycoConfig
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

func NewEpaycoService(config core.EpaycoConfig) EpaycoService {
	return &epaycoService{config}
}
