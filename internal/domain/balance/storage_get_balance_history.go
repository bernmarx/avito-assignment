package balance

import (
	"os"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func (s *Storage) GetBalanceHistory(account_id int, balance_id int, sort string, page int64) ([]Transaction, error) {
	account_owns_balance, err := s.CheckAccountBalanceOwnership(account_id, balance_id)

	if err != nil {
		return nil, errors.New(err.Error(), 500)
	}

	if !account_owns_balance {
		return nil, errors.New("accound_id does not own balance_id", 200)
	}

	tx, err := s.database.Begin()
	if err != nil {
		return nil, errors.New(err.Error(), 503)
	}

	defer tx.Rollback()

	balance_history_select_query := sq.
		Select("operation::text, created_at, value::numeric::float8, receiver_account_id, sender_account_id").
		From("balance_history").
		Where(sq.Eq{"balance_id": balance_id})

	switch sort {
	case "date_asc":
		balance_history_select_query = balance_history_select_query.OrderBy("created_at ASC")
	case "date_desc":
		balance_history_select_query = balance_history_select_query.OrderBy("created_at DESC")
	case "value_asc":
		balance_history_select_query = balance_history_select_query.OrderBy("value ASC")
	case "value_desc":
		balance_history_select_query = balance_history_select_query.OrderBy("value DESC")
	}

	if page > 0 {
		limit, err := strconv.ParseInt(os.Getenv("MAX_HISTORY_PAGE_LEN"), 0, 0)
		if err != nil {
			return nil, errors.New(err.Error(), 500)
		}

		balance_history_select_query = balance_history_select_query.Limit(uint64(limit)).Offset(uint64(limit * (page - 1)))
	}

	balance_history_select, args, err := balance_history_select_query.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return nil, errors.New(err.Error(), 500)
	}

	rows, err := tx.Query(balance_history_select, args...)
	if err != nil {
		return nil, errors.New(err.Error(), 503)
	}

	var balance_history []Transaction

	for rows.Next() {
		var transaction Transaction
		err = rows.Scan(&transaction.Operation,
			&transaction.CreatedAt,
			&transaction.Value,
			&transaction.ReceiverAccountID,
			&transaction.SenderAccountID)

		if err != nil {
			return nil, errors.New(err.Error(), 500)
		}

		balance_history = append(balance_history, transaction)
	}

	return balance_history, tx.Commit()
}
