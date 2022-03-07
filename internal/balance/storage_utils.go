package balance

func GetSqlStmtsPage(sort string) []string {
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
