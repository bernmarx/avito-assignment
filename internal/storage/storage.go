package storage

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
	getTransfer := `SELECT transfer_time, receiver_id, amount::numeric::FLOAT8 FROM transfer_journal WHERE sender_id = $1`

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

	transfers, err := s.Query(getTransfer, id)
	if err != nil {
		return t, err
	}

	for transfers.Next() {
		var receiver int
		var time string
		var amount float32
		err := transfers.Scan(&receiver, &time, &amount)
		if err != nil {
			return t, err
		}

		t.Th = append(t.Th, TransferHistory{Receiver: receiver, Time: time, Amount: amount})
	}

	return t, nil
}
