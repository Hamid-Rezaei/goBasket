package model

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty"`
	Basket    Basket `json:"basket,omitempty"`
}
