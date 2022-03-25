package balance

import (
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func (s *Storage) WithdrawMoney(account_id int, balance_id int, amount float32) error {

	account_owns_balance, err := s.CheckAccountBalanceOwnership(account_id, balance_id)

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	if !account_owns_balance {
		return errors.New("account_id does not own balance_id", 200)
	}

	balance_enough_money, err := s.CheckBalanceEnoughMoney(balance_id, amount)

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	if !balance_enough_money {
		return errors.New("not enough money on balance_id", 200)
	}

	tx, err := s.database.Begin()
	if err != nil {
		return errors.New(err.Error(), 503)
	}

	defer tx.Rollback()

	update_balance, args, err := sq.
		Update("balance").
		Set("balance", sq.Expr("balance - ?::float8::numeric::money", amount)).
		Where(sq.Eq{"id": balance_id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(update_balance, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	balance_history_insert, args, err := sq.
		Insert("balance_history").
		Columns("balance_id, operation, created_at, value").
		Values(balance_id, "withdraw", time.Now().UTC(), amount).
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
