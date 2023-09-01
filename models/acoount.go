package models

type Account struct {
	ID             uint64  `json:"id,omitempty"`
	Username       string  `json:"username,omitempty"`
	CurrentBalance float64 `json:"current_balance,omitempty"`
}
