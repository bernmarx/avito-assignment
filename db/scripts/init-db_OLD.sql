CREATE TABLE IF NOT EXISTS account_balance (
    id INT PRIMARY KEY,
    balance MONEY NOT NULL CONSTRAINT positive_balance CHECK(balance::money::numeric::float8 >= 0)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_account_balance ON account_balance (
    id
);

CREATE OR REPLACE PROCEDURE balance_deposit(_id INT, deposit FLOAT8)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO account_balance (id, balance)
    VALUES (_id, deposit::numeric::money)
    ON CONFLICT (id)
    DO
        UPDATE SET balance = EXCLUDED.balance + account_balance.balance;

    INSERT INTO deposit_journal (account_id, amount, deposit_time)
    VALUES (_id, deposit::numeric::money, now());
COMMIT;
END;
$$;

CREATE OR REPLACE PROCEDURE balance_withdraw(_id INT, withdraw FLOAT8)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE account_balance
        SET balance = balance - withdraw::numeric::money
        WHERE id = _id;
    IF NOT FOUND THEN RAISE EXCEPTION 'user was not found';
    END IF;

    INSERT INTO withdraw_journal (account_id, amount, withdraw_time)
    VALUES (_id, withdraw::numeric::money, now());
COMMIT;
END;
$$;

CREATE OR REPLACE PROCEDURE balance_transfer(sender INT, receiver INT, amount FLOAT8)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE account_balance
        SET balance = balance - amount::numeric::money
        WHERE id = sender;
    IF NOT FOUND THEN RAISE EXCEPTION 'user was not found';
    END IF;

    UPDATE account_balance
        SET balance = balance + amount::numeric::money
        WHERE id = receiver;
    IF NOT FOUND THEN RAISE EXCEPTION 'user was not found';
    END IF;

    INSERT INTO transfer_journal (sender_id, receiver_id, amount, transfer_time)
    VALUES (sender, receiver, amount::numeric::money, now());
COMMIT;
END;
$$;

CREATE OR REPLACE FUNCTION balance_get(_id INT)
RETURNS FLOAT8 AS $bal$
DECLARE
    bal MONEY;
BEGIN
    SELECT balance INTO bal FROM account_balance WHERE id = _id;
    IF NOT FOUND THEN RETURN -1.0;
    END IF;
    RETURN bal::numeric::float8;
END; $bal$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS deposit_journal (
    deposit_id SERIAL PRIMARY KEY,
    account_id INT,
    amount MONEY NOT NULL CONSTRAINT positive_deposit CHECK(amount::money::numeric::float8 > 0),
    deposit_time TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_deposit_journal ON deposit_journal (
    account_id
);

CREATE TABLE IF NOT EXISTS withdraw_journal (
    withdraw_id SERIAL PRIMARY KEY,
    account_id INT,
    amount MONEY NOT NULL CONSTRAINT positive_withdraw CHECK(amount::money::numeric::float8 > 0),
    withdraw_time TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_withdraw_journal ON withdraw_journal (
    account_id
);

CREATE TABLE IF NOT EXISTS transfer_journal (
    transfer_id SERIAL PRIMARY KEY,
    sender_id INT,
    receiver_id INT,
    amount MONEY NOT NULL CONSTRAINT positive_transfer CHECK(amount::money::numeric::float8 > 0),
    transfer_time TIMESTAMP WITH TIME ZONE
);
