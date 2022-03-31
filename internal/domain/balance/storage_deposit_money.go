package balance

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

// DepositMoney deposits 'amount' to 'balanceID'
func (s *Storage) DepositMoney(accountID int, balanceID int, amount float32) error {
	tx, err := s.database.Begin()
	if err != nil {
		return errors.New(err.Error(), 503)
	}

	defer tx.Rollback()

	accountUpsert, args, err := sq.
		Insert("account").
		Columns("id, created_at").
		Values(accountID, time.Now().UTC()).
		Suffix("ON CONFLICT (id) DO NOTHING").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(accountUpsert, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	balanceUpsert, args, err := sq.
		Insert("balance").
		Columns("id, balance, changed_at").
		Values(balanceID, amount, time.Now().UTC()).
		Suffix("ON CONFLICT (id) DO UPDATE SET balance = EXCLUDED.balance + balance.balance").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(balanceUpsert, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	accountBalanceUpsert, args, err := sq.
		Insert("account_balance").
		Columns("account_id, balance_id").
		Values(accountID, balanceID).
		Suffix("ON CONFLICT DO NOTHING").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(accountBalanceUpsert, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	balanceHistoryInsert, args, err := sq.
		Insert("balance_history").
		Columns("balance_id, operation, created_at, value").
		Values(balanceID, "deposit", time.Now().UTC(), amount).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(balanceHistoryInsert, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	return tx.Commit()
}
