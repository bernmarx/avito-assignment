package balance

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func (s *Storage) CheckAccountBalanceOwnership(account_id int, balance_id int) (bool, error) {

	check_account_balance_ownership, args, err := sq.
		Select("*").
		From("account_balance").
		Where(sq.Eq{"account_id": account_id, "balance_id": balance_id}).
		Prefix("SELECT EXISTS(").Suffix(")").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return false, errors.New(err.Error(), 500)
	}

	var account_owns_balance bool

	err = s.database.QueryRow(check_account_balance_ownership, args...).Scan(&account_owns_balance)

	return account_owns_balance, err
}
