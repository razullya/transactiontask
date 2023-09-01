package storage

import (
	"database/sql"
	"fmt"
	"task/models"
)

type Account interface {
	CreateAccount(username string, balance float64) error
	GetAllAccounts() ([]models.Account, error)
	GetTransactionByAccountID(id int) ([]models.Transaction, error)
	GetAccountByID(id int) (models.Account, error)
	GetBalanceByID(id int) (int64, error)
	UpdateBalance(balance int64, id int64) error
}

type AccountStorage struct {
	db *sql.DB
}

func newAccountStorage(db *sql.DB) *AccountStorage {
	return &AccountStorage{
		db: db,
	}
}

func (a *AccountStorage) CreateAccount(username string, balance float64) error {
	_, err := a.db.Exec("INSERT INTO account (username, balance) VALUES ($1, $2)", username, balance)
	if err != nil {
		return fmt.Errorf("CreateAccount: %s", err.Error())
	}
	return nil
}

func (a *AccountStorage) GetAllAccounts() ([]models.Account, error) {
	var accs []models.Account
	rows, err := a.db.Query(`SELECT id, username, balance FROM account;`)
	if err != nil {
		return nil, fmt.Errorf("CreateAccount: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var account models.Account
		err := rows.Scan(&account.ID, &account.Username, &account.CurrentBalance)
		if err != nil {
			return nil, fmt.Errorf("CreateAccount: %s", err.Error())
		}
		accs = append(accs, account)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("CreateAccount: %s", err.Error())
	}
	return accs, nil
}

func (a *AccountStorage) GetTransactionByAccountID(id int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	rows, err := a.db.Query(`
			SELECT id, value, account_id, type_transaction, recepient_id, date, spread_months
			FROM transaction
			WHERE account_id = $1;`,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionByAccountID: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.Value,
			&transaction.AccountID,
			&transaction.TypeTransaction,
			&transaction.RecepientID,
			&transaction.Date,
			&transaction.SpreadMonths,
		)
		if err != nil {
			return nil, fmt.Errorf("GetTransactionByAccountID: %s", err.Error())
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("GetTransactionByAccountID: %s", err.Error())
	}
	return transactions, nil
}

func (a *AccountStorage) GetAccountByID(id int) (models.Account, error) {
	var acc models.Account

	if err := a.db.QueryRow(`
			SELECT id, username, balance
			FROM account
			WHERE id = $1;`,
		id,
	).Scan(&acc.ID, &acc.Username, &acc.CurrentBalance); err != nil {
		return models.Account{}, fmt.Errorf("GetAccountByID: %s", err.Error())
	}
	return acc, nil
}

func (a *AccountStorage) GetBalanceByID(id int) (int64, error) {
	var balance int64

	if err := a.db.QueryRow(`
			SELECT balance
			FROM account
			WHERE id = $1;`,
		id,
	).Scan(&balance); err != nil {
		return 0, fmt.Errorf("GetBalanceByID: %s", err.Error())
	}
	return balance, nil
}

func (a *AccountStorage) UpdateBalance(balance int64, id int64) error {
	_, err := a.db.Exec(`UPDATE account SET balance = balance + $1 WHERE id = $2`, balance, id)
	if err != nil {
		return fmt.Errorf("UpdateBalance: %s", err.Error())
	}
	return nil
}
