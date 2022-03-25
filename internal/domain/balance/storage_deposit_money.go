package balance

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func (s *Storage) DepositMoney(account_id int, balance_id int, amount float32) error {
	tx, err := s.database.Begin()
	if err != nil {
		return errors.New(err.Error(), 503)
	}

	defer tx.Rollback()

	account_upsert, args, err := sq.
		Insert("account").
		Columns("id, created_at").
		Values(account_id, time.Now().UTC()).
		Suffix("ON CONFLICT (id) DO NOTHING").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(account_upsert, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	balance_upsert, args, err := sq.
		Insert("balance").
		Columns("id, balance, changed_at").
		Values(balance_id, amount, time.Now().UTC()).
		Suffix("ON CONFLICT (id) DO UPDATE SET balance = EXCLUDED.balance + balance.balance").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(balance_upsert, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	account_balance_upsert, args, err := sq.
		Insert("account_balance").
		Columns("account_id, balance_id").
		Values(account_id, balance_id).
		Suffix("ON CONFLICT DO NOTHING").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(account_balance_upsert, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	balance_history_insert, args, err := sq.
		Insert("balance_history").
		Columns("balance_id, operation, created_at, value").
		Values(balance_id, "deposit", time.Now().UTC(), amount).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(balance_history_insert, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	return tx.Commit()
}
