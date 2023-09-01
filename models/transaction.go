package models

import "time"

type Transaction struct {
	ID              uint64    `json:"id,omitempty"`
	Value           float64   `json:"value,omitempty"`
	AccountID       uint64    `json:"account_id,omitempty"`
	TypeTransaction uint8     `json:"type_transaction,omitempty"`
	RecepientID     uint64    `json:"recepient_id,omitempty"`
	Date            time.Time `json:"date,omitempty"`
	SpreadMonths    float64   `json:"spread_months,omitempty"`
}
