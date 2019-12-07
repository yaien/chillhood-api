package models

type User struct {
	Name     string
	Email    string
	Password string
	Role     string
}

type Guess struct {
	ID       string
	Cart     *Cart
	Shipping *Shipping
}
