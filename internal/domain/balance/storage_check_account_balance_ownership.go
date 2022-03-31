package balance

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func (s *Storage) CheckAccountBalanceOwnership(accountId int, balanceId int) (bool, error) {

	checkAccountBalanceOwnership, args, err := sq.
		Select("*").
		From("account_balance").
		Where(sq.Eq{"account_id": accountId, "balance_id": balanceId}).
		Prefix("SELECT EXISTS(").Suffix(")").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return false, errors.New(err.Error(), 500)
	}

	var accountOwnsBalance bool

	err = s.database.QueryRow(checkAccountBalanceOwnership, args...).Scan(&accountOwnsBalance)

	return accountOwnsBalance, err
}
