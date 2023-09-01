package storage

import (
	"database/sql"
	"fmt"
	"task/models"
)

type Transaction interface {
	CreateTransaction(transaction models.Transaction) error
	GetTransactionByID(id int) (models.Transaction, error)
	GetAllTransactions() ([]models.Transaction, error)
}

type TransactionStorage struct {
	db *sql.DB
}

func newTransactionStorage(db *sql.DB) *TransactionStorage {
	return &TransactionStorage{
		db: db,
	}
}

func (t *TransactionStorage) CreateTransaction(transaction models.Transaction) error {
	_, err := t.db.Exec(`
			INSERT INTO transaction (value, account_id, type_transaction, date, recepient_id, spread_months)
			VALUES ($1, $2, $3, CURRENT_TIMESTAMP, $4, $5);`,
		transaction.Value,
		transaction.AccountID,
		transaction.TypeTransaction,
		transaction.RecepientID,
		transaction.SpreadMonths,
	)
	if err != nil {
		return fmt.Errorf("CreateTransaction: %s", err.Error())
	}
	return nil
}

func (t *TransactionStorage) GetTransactionByID(id int) (models.Transaction, error) {
	var transaction models.Transaction
	if err := t.db.QueryRow(`
			SELECT id, value, account_id, type_transaction, recepient_id, date, spread_months
			FROM transaction
			WHERE id = $1;
	`, id).Scan(
		&transaction.ID,
		&transaction.Value,
		&transaction.AccountID,
		&transaction.TypeTransaction,
		&transaction.RecepientID,
		&transaction.Date,
		&transaction.SpreadMonths,
	); err != nil {
		return models.Transaction{}, fmt.Errorf("GetTransactionByID: %s", err.Error())
	}
	return transaction, nil
}

func (t *TransactionStorage) GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction

	rows, err := t.db.Query(
		`
		SELECT id, value, account_id, type_transaction, recepient_id, date, spread_months
		FROM transaction;
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("GetAllTransactions: %s", err.Error())
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
			return nil, fmt.Errorf("GetAllTransactions: %s", err.Error())
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllTransactions: %s", err.Error())
	}
	return transactions, nil
}
