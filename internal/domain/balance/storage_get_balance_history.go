package balance

import (
	"encoding/json"
	"os"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func (s *Storage) GetBalanceHistory(accountId int, balanceId int, sort string, page int64) ([]byte, error) {
	accountOwnsBalance, err := s.CheckAccountBalanceOwnership(accountId, balanceId)

	if err != nil {
		return nil, errors.New(err.Error(), 500)
	}

	if !accountOwnsBalance {
		return nil, errors.New("account_id does not own balance_id", 200)
	}

	tx, err := s.database.Begin()
	if err != nil {
		return nil, errors.New(err.Error(), 503)
	}

	defer tx.Rollback()

	balanceHistorySelectQuery := sq.
		Select("operation::text, created_at, value::numeric::float8, receiver_account_id, sender_account_id").
		From("balance_history").
		Where(sq.Eq{"balance_id": balanceId})

	switch sort {
	case "date_asc":
		balanceHistorySelectQuery = balanceHistorySelectQuery.OrderBy("created_at ASC")
	case "date_desc":
		balanceHistorySelectQuery = balanceHistorySelectQuery.OrderBy("created_at DESC")
	case "value_asc":
		balanceHistorySelectQuery = balanceHistorySelectQuery.OrderBy("value ASC")
	case "value_desc":
		balanceHistorySelectQuery = balanceHistorySelectQuery.OrderBy("value DESC")
	}

	if page > 0 {
		limit, err := strconv.ParseInt(os.Getenv("MAX_HISTORY_PAGE_LEN"), 0, 0)
		if err != nil {
			return nil, errors.New(err.Error(), 500)
		}

		balanceHistorySelectQuery = balanceHistorySelectQuery.Limit(uint64(limit)).Offset(uint64(limit * (page - 1)))
	}

	balanceHistorySelect, args, err := balanceHistorySelectQuery.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return nil, errors.New(err.Error(), 500)
	}

	rows, err := tx.Query(balanceHistorySelect, args...)
	if err != nil {
		return nil, errors.New(err.Error(), 503)
	}

	var balanceHistory []Transaction

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

		balanceHistory = append(balanceHistory, transaction)
	}

	j, err := json.Marshal(balanceHistory)
	if err != nil {
		return nil, err
	}

	return j, tx.Commit()
}
