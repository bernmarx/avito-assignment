package balance

import (
	"database/sql"

	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func getSqlStmtsPage(sort string) []string {
	stmts := make([]string, 4)
	switch sort {
	case "amount_asc":
		stmts[0] = `SELECT deposit_time, amount::numeric::FLOAT8
						FROM deposit_journal WHERE account_id = $1 ORDER BY amount ASC LIMIT $2 OFFSET $3`
		stmts[1] = `SELECT withdraw_time, amount::numeric::FLOAT8
						FROM withdraw_journal WHERE account_id = $1 ORDER BY amount ASC LIMIT $2 OFFSET $3`
		stmts[2] = `SELECT transfer_time, receiver_id, amount::numeric::FLOAT8
						FROM transfer_journal WHERE sender_id = $1 ORDER BY amount ASC LIMIT $2 OFFSET $3`
		stmts[3] = `SELECT transfer_time, sender_id, amount::numeric::FLOAT8
						FROM transfer_journal WHERE receiver_id = $1 ORDER BY amount ASC LIMIT $2 OFFSET $3`
		return stmts
	case "amount_desc":
		stmts[0] = `SELECT deposit_time, amount::numeric::FLOAT8
						FROM deposit_journal WHERE account_id = $1 ORDER BY amount DESC LIMIT $2 OFFSET $3`
		stmts[1] = `SELECT withdraw_time, amount::numeric::FLOAT8
						FROM withdraw_journal WHERE account_id = $1 ORDER BY amount DESC LIMIT $2 OFFSET $3`
		stmts[2] = `SELECT transfer_time, receiver_id, amount::numeric::FLOAT8
						FROM transfer_journal WHERE sender_id = $1 ORDER BY amount DESC LIMIT $2 OFFSET $3`
		stmts[3] = `SELECT transfer_time, sender_id, amount::numeric::FLOAT8
						FROM transfer_journal WHERE receiver_id = $1 ORDER BY amount DESC LIMIT $2 OFFSET $3`
		return stmts
	case "date_asc":
		stmts[0] = `SELECT deposit_time, amount::numeric::FLOAT8
						FROM deposit_journal WHERE account_id = $1 ORDER BY deposit_time ASC LIMIT $2 OFFSET $3`
		stmts[1] = `SELECT withdraw_time, amount::numeric::FLOAT8
						FROM withdraw_journal WHERE account_id = $1 ORDER BY withdraw_time ASC LIMIT $2 OFFSET $3`
		stmts[2] = `SELECT transfer_time, receiver_id, amount::numeric::FLOAT8
						FROM transfer_journal WHERE sender_id = $1 ORDER BY transfer_time ASC LIMIT $2 OFFSET $3`
		stmts[3] = `SELECT transfer_time, sender_id, amount::numeric::FLOAT8
						FROM transfer_journal WHERE receiver_id = $1 ORDER BY transfer_time ASC LIMIT $2 OFFSET $3`
	case "date_desc":
		stmts[0] = `SELECT deposit_time, amount::numeric::FLOAT8
						FROM deposit_journal WHERE account_id = $1 ORDER BY deposit_time DESC LIMIT $2 OFFSET $3`
		stmts[1] = `SELECT withdraw_time, amount::numeric::FLOAT8
						FROM withdraw_journal WHERE account_id = $1 ORDER BY withdraw_time DESC LIMIT $2 OFFSET $3`
		stmts[2] = `SELECT transfer_time, receiver_id, amount::numeric::FLOAT8
						FROM transfer_journal WHERE sender_id = $1 ORDER BY transfer_time DESC LIMIT $2 OFFSET $3`
		stmts[3] = `SELECT transfer_time, sender_id, amount::numeric::FLOAT8
						FROM transfer_journal WHERE receiver_id = $1 ORDER BY transfer_time DESC LIMIT $2 OFFSET $3`
	}
	return stmts
}

func getDepositHistory(d *sql.Rows) ([]DepositHistory, error) {
	var dh []DepositHistory

	for d.Next() {
		var time string
		var amount float32
		err := d.Scan(&time, &amount)
		if err != nil {
			return dh, errors.New(err.Error(), 500)
		}

		dh = append(dh, DepositHistory{Time: time, Amount: amount})
	}

	return dh, nil
}

func getWithdrawalHistory(w *sql.Rows) ([]WithdrawalHistory, error) {
	var wh []WithdrawalHistory

	for w.Next() {
		var time string
		var amount float32
		err := w.Scan(&time, &amount)
		if err != nil {
			return wh, errors.New(err.Error(), 500)
		}

		wh = append(wh, WithdrawalHistory{Time: time, Amount: amount})
	}

	return wh, nil
}

func getSendsHistory(s *sql.Rows) ([]SendHistory, error) {
	var sh []SendHistory

	for s.Next() {
		var time string
		var receiver int
		var amount float32

		err := s.Scan(&time, &receiver, &amount)
		if err != nil {
			return sh, errors.New(err.Error(), 500)
		}

		sh = append(sh, SendHistory{ReceiverID: receiver, Time: time, Amount: amount})
	}

	return sh, nil
}

func getReceivesHistory(r *sql.Rows) ([]ReceiveHistory, error) {
	var rh []ReceiveHistory

	for r.Next() {
		var time string
		var sender int
		var amount float32

		err := r.Scan(&time, &sender, &amount)
		if err != nil {
			return rh, errors.New(err.Error(), 500)
		}

		rh = append(rh, ReceiveHistory{SenderID: sender, Time: time, Amount: amount})
	}

	return rh, nil
}
