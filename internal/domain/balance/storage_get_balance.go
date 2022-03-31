package balance

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

// GetBalance returns balance of 'balanceID'
func (s *Storage) GetBalance(accountID int, balanceID int) (float32, error) {
	accountOwnsBalance, err := s.CheckAccountBalanceOwnership(accountID, balanceID)

	if err != nil {
		return 0.0, errors.New(err.Error(), 500)
	}

	if !accountOwnsBalance {
		return 0.0, errors.New("account_id does not own balance_id", 400)
	}

	tx, err := s.database.Begin()
	if err != nil {
		return 0.0, errors.New(err.Error(), 503)
	}

	defer tx.Rollback()

	balanceSelect, args, err := sq.
		Select("balance::numeric::float8").
		From("balance").
		Where(sq.Eq{"id": balanceID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0.0, errors.New(err.Error(), 500)
	}

	var balance float32

	err = tx.QueryRow(balanceSelect, args...).Scan(&balance)
	if err != nil {
		return 0.0, errors.New(err.Error(), 500)
	}

	return balance, nil
}
