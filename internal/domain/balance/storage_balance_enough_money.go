package balance

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func (s *Storage) CheckBalanceEnoughMoney(balanceId int, amount float32) (bool, error) {
	checkBalanceEnoughMoney, args, err := sq.Select("*").
		From("balance").
		Where(sq.Eq{"id": balanceId}, sq.Expr("balance >= ?::float8::numeric::money", amount)).
		Prefix("SELECT EXISTS(").Suffix(")").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return false, errors.New(err.Error(), 500)
	}

	var balanceEnoughMoney bool

	err = s.database.QueryRow(checkBalanceEnoughMoney, args...).Scan(&balanceEnoughMoney)

	return balanceEnoughMoney, err
}
