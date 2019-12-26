package services

import "testing"

import "net/url"

import "github.com/yaien/clothes-store-api/core"

func TestCheckoutArgs(t *testing.T) {
	testcases := []struct {
		baseURL      string
		response     string
		confirmation string
	}{
		{
			baseURL:      "https://store.com",
			response:     "https://store.com/api/v1/epayco/response",
			confirmation: "https://store.com/api/v1/epayco/confirmation",
		},
		{
			baseURL:      "http://localhost:4200",
			response:     "http://localhost:4200/api/v1/epayco/response",
			confirmation: "http://localhost:4200/api/v1/epayco/confirmation",
		},
	}
	for _, testcase := range testcases {
		baseURL, err := url.Parse(testcase.baseURL)
		if err != nil {
			t.Error(err)
			continue
		}
		service := &epaycoService{baseURL: baseURL, config: &core.EpaycoConfig{}}
		args := service.CheckoutArgs()

		if args.Response != testcase.response {
			t.Errorf("Expected response url to be '%s', received '%s'",
				testcase.response, args.Response)
		}
		if args.Confirmation != testcase.confirmation {
			t.Errorf("Expected confirmation url to be '%s', received: '%s'",
				testcase.confirmation, args.Confirmation)
		}
	}
}
