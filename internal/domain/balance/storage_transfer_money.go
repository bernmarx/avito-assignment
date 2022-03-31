package balance

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func (s *Storage) TransferMoney(senderAccountId int, senderBalanceId int,
	receiverAccountId int, receiverBalanceId int, amount float32) error {

	accountOwnsBalance, err := s.CheckAccountBalanceOwnership(senderAccountId, senderBalanceId)

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	if !accountOwnsBalance {
		return errors.New("sender_account_id does not own sender_balance_id", 200)
	}

	accountOwnsBalance, err = s.CheckAccountBalanceOwnership(receiverAccountId, receiverBalanceId)

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	if !accountOwnsBalance {
		return errors.New("receiver_account_id does not own receiver_balance_id", 200)
	}

	senderEnoughMoney, err := s.CheckBalanceEnoughMoney(senderBalanceId, amount)

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	if !senderEnoughMoney {
		return errors.New("not enough money on sender_balance_id", 200)
	}

	tx, err := s.database.Begin()
	if err != nil {
		return errors.New(err.Error(), 503)
	}

	defer tx.Rollback()

	updateSenderBalance, args, err := sq.
		Update("balance").
		Set("balance", sq.Expr("balance - ?::float8::numeric::money", amount)).
		Where(sq.Eq{"id": senderBalanceId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(updateSenderBalance, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	updateReceiverBalance, args, err := sq.
		Update("balance").
		Set("balance", sq.Expr("balance + ?::float8::numeric::money", amount)).
		Where(sq.Eq{"id": receiverBalanceId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(updateReceiverBalance, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	senderBalanceHistoryInsert, args, err := sq.
		Insert("balance_history").
		Columns("balance_id, operation, created_at, value, receiver_account_id, sender_account_id").
		Values(senderBalanceId, "transfer", time.Now().UTC(), amount, receiverAccountId, senderAccountId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(senderBalanceHistoryInsert, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	receiverBalanceHistoryInsert, args, err := sq.
		Insert("balance_history").
		Columns("balance_id, operation, created_at, value, receiver_account_id, sender_account_id").
		Values(receiverBalanceId, "transfer", time.Now().UTC(), amount, receiverAccountId, senderAccountId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(receiverBalanceHistoryInsert, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	return tx.Commit()
}
