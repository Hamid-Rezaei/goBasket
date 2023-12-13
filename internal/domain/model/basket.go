package model

type Basket struct {
	ID    uint64 `json:"id,omitempty"`
	Data  string `json:"data,omitempty"`
	State string `json:"state,omitempty"`
}
