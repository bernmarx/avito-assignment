CREATE TYPE operation_type AS ENUM ('deposit', 'withdraw', 'transfer');

CREATE TABLE IF NOT EXISTS account (
    id INT PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS balance (
    id INT PRIMARY KEY,
    balance MONEY NOT NULL CONSTRAINT non_negative_balance CHECK(balance::money::numeric::float8 >= 0),
    changed_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS account_balance (
    account_id INT REFERENCES account(id) ON DELETE NO ACTION,
    balance_id INT REFERENCES balance(id) ON DELETE NO ACTION
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_account_balance ON account_balance (
    account_id, balance_id
);

CREATE TABLE IF NOT EXISTS balance_history (
    id SERIAL PRIMARY KEY,
    balance_id INT REFERENCES balance(id),
    operation operation_type,
    created_at TIMESTAMP WITH TIME ZONE,
    value MONEY NOT NULL CONSTRAINT positive_value CHECK(value::money::numeric::float8 > 0),
    receiver_account_id INT,
    sender_account_id INT
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_balance_history ON balance_history (
    balance_id, created_at, receiver_account_id, sender_account_id
);
