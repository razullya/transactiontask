package service

import (
	"strconv"
	"task/internal/storage"
	"task/models"
)

type Account interface {
	CreateAccount(username string, balance float64) error
	GetAllAccounts() ([]models.Account, error)
	GetTransactionByAccountID(id string) ([]models.Transaction, error)
	GetAccountByID(id string) (models.Account, error)
}

type AccountService struct {
	storage storage.Storage
}

func newAccountService(storage *storage.Storage) *AccountService {
	return &AccountService{
		storage: *storage,
	}
}

func (a *AccountService) GetAccountByAccountName(id string) (models.Account, error) {
	var Account models.Account

	return Account, nil
}

func (a *AccountService) CreateAccount(username string, balance float64) error {
	return a.storage.CreateAccount(username, balance)
}

func (a *AccountService) GetAllAccounts() ([]models.Account, error) {
	return a.storage.GetAllAccounts()
}

func (a *AccountService) GetTransactionByAccountID(id string) ([]models.Transaction, error) {
	idd, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return a.storage.GetTransactionByAccountID(idd)
}

func (a *AccountService) GetAccountByID(id string) (models.Account, error) {
	idd, err := strconv.Atoi(id)
	if err != nil {
		return models.Account{}, err
	}
	return a.storage.GetAccountByID(idd)
}
