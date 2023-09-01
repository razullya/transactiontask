package storage

import (
	"database/sql"
)

type Storage struct {
	Account
	Transaction
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Account:     newAccountStorage(db),
		Transaction: newTransactionStorage(db),
	}
}
