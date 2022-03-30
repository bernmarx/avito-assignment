package balance

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func (s *Storage) GetBalance(account_id int, balance_id int) (float32, error) {
	account_owns_balance, err := s.CheckAccountBalanceOwnership(account_id, balance_id)

	if err != nil {
		return 0.0, errors.New(err.Error(), 500)
	}

	if !account_owns_balance {
		return 0.0, errors.New("accound_id does not own balance_id", 400)
	}

	tx, err := s.database.Begin()
	if err != nil {
		return 0.0, errors.New(err.Error(), 503)
	}

	defer tx.Rollback()

	balance_select, args, err := sq.
		Select("balance::numeric::float8").
		From("balance").
		Where(sq.Eq{"id": balance_id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0.0, errors.New(err.Error(), 500)
	}

	var balance float32

	err = tx.QueryRow(balance_select, args...).Scan(&balance)
	if err != nil {
		return 0.0, errors.New(err.Error(), 500)
	}

	return balance, nil
}
