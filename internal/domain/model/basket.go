package model

type Basket struct {
	Data   string `json:"data,omitempty"`
	State  string `json:"state,omitempty"`
	UserID uint   `json:"user_id,omitempty"`
}
