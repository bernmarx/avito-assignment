package balance

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func (s *Storage) CheckBalanceEnoughMoney(balance_id int, amount float32) (bool, error) {
	check_balance_enough_money, args, err := sq.Select("*").
		From("balance").
		Where(sq.Eq{"id": balance_id}, sq.Expr("balance >= ?::float8::numeric::money", amount)).
		Prefix("SELECT EXISTS(").Suffix(")").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return false, errors.New(err.Error(), 500)
	}

	var balance_enough_money bool

	err = s.database.QueryRow(check_balance_enough_money, args...).Scan(&balance_enough_money)

	return balance_enough_money, err
}
