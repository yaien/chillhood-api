package entity

type ShippingStatus string

const (
	Preparing = ShippingStatus("preparing")
	Sended    = ShippingStatus("sended")
)

type Shipping struct {
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Phone      string         `json:"phone"`
	Address    string         `json:"address"`
	City       string         `json:"city"`
	Province   string         `json:"province"`
	Country    string         `json:"country"`
	PostalCode string         `json:"postalCode"`
	Status     ShippingStatus `json:"status,omitempty"`
	Transport  *Transport     `json:"transport,omitempty"`
}

type Transport struct {
	Provider string `json:"provider"`
	Guide    string `json:"guide"`
}
