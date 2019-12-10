package models

type Shipping struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
	Status     string `json:"status"`
}
