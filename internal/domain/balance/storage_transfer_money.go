package balance

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func (s *Storage) TransferMoney(sender_account_id int, sender_balance_id int,
	receiver_account_id int, receiver_balance_id int, amount float32) error {

	account_owns_balance, err := s.CheckAccountBalanceOwnership(sender_account_id, sender_balance_id)

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	if !account_owns_balance {
		return errors.New("sender_accound_id does not own sender_balance_id", 200)
	}

	account_owns_balance, err = s.CheckAccountBalanceOwnership(receiver_account_id, receiver_balance_id)

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	if !account_owns_balance {
		return errors.New("receiver_account_id does not own receiver_balance_id", 200)
	}

	sender_enough_money, err := s.CheckBalanceEnoughMoney(sender_balance_id, amount)

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	if !sender_enough_money {
		return errors.New("not enough money on sender_balance_id", 200)
	}

	tx, err := s.database.Begin()
	if err != nil {
		return errors.New(err.Error(), 503)
	}

	defer tx.Rollback()

	update_sender_balance, args, err := sq.
		Update("balance").
		Set("balance", sq.Expr("balance - ?::float8::numeric::money", amount)).
		Where(sq.Eq{"id": sender_balance_id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(update_sender_balance, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	update_receiver_balance, args, err := sq.
		Update("balance").
		Set("balance", sq.Expr("balance + ?::float8::numeric::money", amount)).
		Where(sq.Eq{"id": receiver_balance_id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(update_receiver_balance, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	sender_balance_history_insert, args, err := sq.
		Insert("balance_history").
		Columns("balance_id, operation, created_at, value, receiver_account_id, sender_account_id").
		Values(sender_balance_id, "transfer", time.Now().UTC(), amount, receiver_account_id, sender_account_id).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(sender_balance_history_insert, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	receiver_balance_history_insert, args, err := sq.
		Insert("balance_history").
		Columns("balance_id, operation, created_at, value, receiver_account_id, sender_account_id").
		Values(receiver_balance_id, "transfer", time.Now().UTC(), amount, receiver_account_id, sender_account_id).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.New(err.Error(), 500)
	}

	_, err = tx.Exec(receiver_balance_history_insert, args...)
	if err != nil {
		return errors.New(err.Error(), 500)
	}

	return tx.Commit()
}
