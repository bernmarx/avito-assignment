package balance

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

// TransferMoney transfers 'amount' from 'senderBalanceID' to 'receiverBalanceID'
func (s *Storage) TransferMoney(senderAccountID int, senderBalanceID int,
	receiverAccountID int, receiverBalanceID int, amount float32) error {

	accountOwnsBalance, err := s.CheckAccountBalanceOwnership(senderAccountID, senderBalanceID)

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	if !accountOwnsBalance {
		return errors.New("sender_account_id does not own sender_balance_id", 200)
	}

	accountOwnsBalance, err = s.CheckAccountBalanceOwnership(receiverAccountID, receiverBalanceID)

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	if !accountOwnsBalance {
		return errors.New("receiver_account_id does not own receiver_balance_id", 200)
	}

	senderEnoughMoney, err := s.CheckBalanceEnoughMoney(senderBalanceID, amount)

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
		Set("balance", sq.Expr("balance - ?::numeric", amount)).
		Where(sq.Eq{"id": senderBalanceID}).
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
		Set("balance", sq.Expr("balance + ?::numeric", amount)).
		Where(sq.Eq{"id": receiverBalanceID}).
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
		Values(senderBalanceID, "transfer", time.Now().UTC(), amount, receiverAccountID, senderAccountID).
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
		Values(receiverBalanceID, "transfer", time.Now().UTC(), amount, receiverAccountID, senderAccountID).
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
