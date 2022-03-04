package storage

import (
	"errors"
	"os"
	"strconv"
)

func (s *Storage) Deposit(id int, amount float32) error {
	sqlstmt := `call balance_deposit($1, $2)`
	_, err := s.Exec(sqlstmt, id, amount)
	return err
}

func (s *Storage) Withdraw(id int, amount float32) error {
	sqlstmt := `call balance_withdraw($1, $2)`
	_, err := s.Exec(sqlstmt, id, amount)
	return err
}

func (s *Storage) Transfer(senderID int, receiverID int, amount float32) error {
	sqlstmt := `call balance_transfer($1, $2, $3)`
	_, err := s.Exec(sqlstmt, senderID, receiverID, amount)
	return err
}

func (s *Storage) GetBalance(id int) (float32, error) {
	var balance float32
	sqlstmt := `SELECT balance_get($1)`
	row := s.QueryRow(sqlstmt, id)
	err := row.Scan(&balance)
	return balance, err
}

func (s *Storage) GetTransactionHistory(id int) (TransactionHistory, error) {
	getDeposit := `SELECT deposit_time, amount::numeric::FLOAT8 FROM deposit_journal WHERE account_id = $1`
	getWithdrawal := `SELECT withdraw_time, amount::numeric::FLOAT8 FROM withdraw_journal WHERE account_id = $1`
	getSend := `SELECT transfer_time, receiver_id, amount::numeric::FLOAT8 FROM transfer_journal WHERE sender_id = $1`
	getReceive := `SELECT transfer_time, sender_id, amount::numeric::FLOAT8 FROM transfer_journal WHERE receiver_id = $1`

	var t TransactionHistory

	deposits, err := s.Query(getDeposit, id)
	if err != nil {
		return t, err
	}

	for deposits.Next() {
		var time string
		var amount float32
		err := deposits.Scan(&time, &amount)
		if err != nil {
			return t, err
		}

		t.Dh = append(t.Dh, DepositHistory{Time: time, Amount: amount})
	}

	withdrawals, err := s.Query(getWithdrawal, id)
	if err != nil {
		return t, err
	}

	for withdrawals.Next() {
		var time string
		var amount float32
		err := withdrawals.Scan(&time, &amount)
		if err != nil {
			return t, err
		}

		t.Wh = append(t.Wh, WithdrawalHistory{Time: time, Amount: amount})
	}

	sends, err := s.Query(getSend, id)
	if err != nil {
		return t, err
	}

	for sends.Next() {
		var receiver int
		var time string
		var amount float32
		err := sends.Scan(&time, &receiver, &amount)
		if err != nil {
			return t, err
		}

		t.Sh = append(t.Sh, SendHistory{ReceiverID: receiver, Time: time, Amount: amount})
	}

	receives, err := s.Query(getReceive, id)
	if err != nil {
		return t, err
	}

	for receives.Next() {
		var sender int
		var time string
		var amount float32
		err := receives.Scan(&time, &sender, &amount)
		if err != nil {
			return t, err
		}

		t.Rh = append(t.Rh, ReceiveHistory{SenderID: sender, Time: time, Amount: amount})
	}

	if t.Dh == nil && t.Sh == nil && t.Wh == nil && t.Rh == nil {
		return t, errors.New("user was not found")
	}

	return t, nil
}

func (s *Storage) GetTransactionHistoryPage(id int, sort string, page int) (TransactionHistory, error) {
	var t TransactionHistory
	stmts := GetSqlStmtsPage(sort)

	pageLength64, err := strconv.ParseInt(os.Getenv("MAX_HISTORY_PAGE_LEN"), 10, 0)
	if err != nil {
		return t, errors.New("could not get MAX_HISTORY_PAGE_LEN")
	}

	pageLength := int(pageLength64)
	offset := pageLength * (page - 1)

	deposits, err := s.Query(stmts[0], id, pageLength, offset)
	if err != nil {
		return t, err
	}

	for deposits.Next() {
		var time string
		var amount float32
		err := deposits.Scan(&time, &amount)
		if err != nil {
			return t, err
		}

		t.Dh = append(t.Dh, DepositHistory{Time: time, Amount: amount})
	}

	withdrawals, err := s.Query(stmts[1], id, pageLength, offset)
	if err != nil {
		return t, err
	}

	for withdrawals.Next() {
		var time string
		var amount float32
		err := withdrawals.Scan(&time, &amount)
		if err != nil {
			return t, err
		}

		t.Wh = append(t.Wh, WithdrawalHistory{Time: time, Amount: amount})
	}

	sends, err := s.Query(stmts[2], id, pageLength, offset)
	if err != nil {
		return t, err
	}

	for sends.Next() {
		var receiver int
		var time string
		var amount float32
		err := sends.Scan(&time, &receiver, &amount)
		if err != nil {
			return t, err
		}

		t.Sh = append(t.Sh, SendHistory{ReceiverID: receiver, Time: time, Amount: amount})
	}

	receives, err := s.Query(stmts[3], id, pageLength, offset)
	if err != nil {
		return t, err
	}

	for receives.Next() {
		var sender int
		var time string
		var amount float32
		err := receives.Scan(&time, &sender, &amount)
		if err != nil {
			return t, err
		}

		t.Rh = append(t.Rh, ReceiveHistory{SenderID: sender, Time: time, Amount: amount})
	}

	if t.Dh == nil && t.Sh == nil && t.Wh == nil {
		return t, errors.New("user was not found or invalid page number")
	}

	return t, nil
}
