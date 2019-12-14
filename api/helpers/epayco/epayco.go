package epayco

// Payment -> epayco payment data
type Payment struct {
	Ref              int          `json:"x_ref_payco"`
	Invoice          string       `json:"x_id_invoice"`
	Description      string       `json:"x_description"`
	Amount           int          `json:"x_amount"`
	AmountContry     int          `json:"x_amount_country"`
	AmountOk         int          `json:"x_amount_ok"`
	AmountBase       int          `json:"x_amount_base"`
	Tax              int          `json:"x_tax"`
	CurrencyCode     string       `json:"x_currency_code"`
	BankName         string       `json:"x_bank_name"`
	Cardnumber       string       `json:"x_cardnumber"`
	Quotas           interface{}  `json:"x_quotas"`
	Response         string       `json:"x_response"`
	ResponseCode     ResponseCode `json:"x_cod_response"`
	ApprovalCode     string       `json:"x_approval_code"`
	TransactionID    string       `json:"x_transaction_id"`
	TransactionDate  string       `json:"x_transaction_date"`
	TransactionState string       `json:"x_transaction_state"`
	Franchise        string       `json:"x_franchise"`
	Test             string       `json:"x_test_request"`
	Signature        string       `json:"x_signature"`
}

type Response struct {
	Success bool     `json:"success"`
	Data    *Payment `json:"data"`
}

type ResponseCode int

const (
	Accepted  ResponseCode = 1
	Rejected  ResponseCode = 2
	Pending   ResponseCode = 3
	Failed    ResponseCode = 4
	Reversed  ResponseCode = 6
	Held      ResponseCode = 7
	Started   ResponseCode = 8
	Expired   ResponseCode = 9
	Abandoned ResponseCode = 10
	Canceled  ResponseCode = 11
	AntiFraud ResponseCode = 12
)
