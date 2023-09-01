package service

import (
	"fmt"
	"strconv"
	"task/internal/storage"
	"task/models"
)

type Transaction interface {
	CreateTransaction(transaction models.Transaction) error
	GetTransactionByID(id string) (models.Transaction, error)
	GetAllTransactions() ([]models.Transaction, error)
}

type TransactionService struct {
	storage storage.Storage
}

func newTransactionService(storage *storage.Storage) *TransactionService {
	return &TransactionService{
		storage: *storage,
	}
}

func (t *TransactionService) CreateTransaction(transaction models.Transaction) error {
	balance, err := t.storage.Account.GetBalanceByID(int(transaction.AccountID))
	if err != nil {
		return err
	}

	if balance-int64(transaction.Value) < 0 {
		return fmt.Errorf("U cannot do that action, p.s: u dont have money")
	}
	//update balance 1
	if err := t.storage.Account.UpdateBalance(int64(transaction.Value*(-1)), int64(transaction.AccountID)); err != nil {
		return err
	}
	//update balance 2
	if err := t.storage.Account.UpdateBalance(int64(transaction.Value), int64(transaction.RecepientID)); err != nil {
		return err
	}
	return t.storage.CreateTransaction(transaction)
}

func (t *TransactionService) GetTransactionByID(id string) (models.Transaction, error) {
	idd, err := strconv.Atoi(id)
	if err != nil {
		return models.Transaction{}, err
	}
	return t.storage.GetTransactionByID(idd)
}

func (t *TransactionService) GetAllTransactions() ([]models.Transaction, error) {
	return t.storage.GetAllTransactions()
}
