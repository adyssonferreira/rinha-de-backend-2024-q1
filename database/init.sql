CREATE TABLE clients (
	id SERIAL PRIMARY KEY,
	nome VARCHAR(50) NOT NULL,
    balance INTEGER NOT NULL DEFAULT 0,
	limite INTEGER NOT NULL
);

CREATE TABLE transactions (
	id SERIAL PRIMARY KEY,
	client_id INTEGER NOT NULL,
	value INTEGER NOT NULL,
	type CHAR(1) NOT NULL,
	description VARCHAR(10) NOT NULL,
	create_at TIMESTAMP NOT NULL DEFAULT NOW(),

	CONSTRAINT fk_clients_transactions_id
		FOREIGN KEY (client_id) REFERENCES clients(id)
);

-- INDEX
CREATE INDEX idx_clients_id ON clients (id);
CREATE INDEX idx_transactions_cliente_id ON transactions (client_id);


--SEED
INSERT INTO clients (nome, limite) VALUES
		('client 1', 1000 * 100),
		('client 2', 800 * 100),
		('client 3', 10000 * 100),
		('client 4', 100000 * 100),
		('client 5', 5000 * 100);