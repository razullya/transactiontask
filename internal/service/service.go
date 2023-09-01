package service

import (
	"task/internal/storage"
)

type Service struct {
	Account
	Transaction
}

func NewService(storages *storage.Storage) *Service {
	return &Service{
		Account:     newAccountService(storages),
		Transaction: newTransactionService(storages),
	}
}
