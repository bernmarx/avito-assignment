package balance

import (
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

// WithdrawMoney withdraws 'amount' from 'balanceID'
func (s *Storage) WithdrawMoney(accountID int, balanceID int, amount float32) error {

	accountOwnsBalance, err := s.CheckAccountBalanceOwnership(accountID, balanceID)

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	if !accountOwnsBalance {
		return errors.New("account_id does not own balance_id", 200)
	}

	balanceEnoughMoney, err := s.CheckBalanceEnoughMoney(balanceID, amount)

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	if !balanceEnoughMoney {
		return errors.New("not enough money on balance_id", 200)
	}

	tx, err := s.database.Begin()
	if err != nil {
		return errors.New(err.Error(), 503)
	}

	defer tx.Rollback()

	updateBalance, args, err := sq.
		Update("balance").
		Set("balance", sq.Expr("balance - ?::float8::numeric::money", amount)).
		Where(sq.Eq{"id": balanceID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(updateBalance, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	balanceHistoryInsert, args, err := sq.
		Insert("balance_history").
		Columns("balance_id, operation, created_at, value").
		Values(balanceID, "withdraw", time.Now().UTC(), amount).
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
