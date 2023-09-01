CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE,
		balance INT
	);

CREATE TABLE IF NOT EXISTS transaction (
		id SERIAL PRIMARY KEY,
		value INT,
		account_id INT,
		type_transaction INT,
		recepient_id INT,
		date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		spread_months DOUBLE PRECISION,
		FOREIGN KEY (account_id) REFERENCES account(id)
	);