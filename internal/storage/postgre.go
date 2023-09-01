package storage

import (
	"database/sql"
	"fmt"
)

const (
	accountTable = `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE,
		balance INT
	);`

	transactionTable = `CREATE TABLE IF NOT EXISTS transaction (
		id SERIAL PRIMARY KEY,
		value INT,
		account_id INT,
		type_transaction INT,
		recepient_id INT,
		date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		spread_months DOUBLE PRECISION,
		FOREIGN KEY (account_id) REFERENCES account(id)
	);`
)

var tables = []string{accountTable, transactionTable}

func InitDB() (*sql.DB, error) {

	db, err := sql.Open("postgres", "host=go_db user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("storage: create db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("storage: ping db: %w", err)
	}
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			return nil, fmt.Errorf("storage: create tables: %w", err)
		}
	}

	return db, nil
}
