package models

type Cart struct {
	Shipping int
	SubTotal int
	Total    int
	Items    []Item
}

type Item struct {
	Name     string
	Price    string
	Quantity int
	Size     string
}
