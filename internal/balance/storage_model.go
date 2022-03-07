//go:generate mockgen -source $GOFILE -destination ./storage_mock.go -package $GOPACKAGE
package balance

import (
	"database/sql"
	"encoding/json"
)

type database interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type Storage struct {
	database
}

func NewStorage(db database) *Storage {
	return &Storage{database: db}
}

type TransactionHistory struct {
	Dh []DepositHistory    `json:"deposit_history"`
	Wh []WithdrawalHistory `json:"withdrawal_history"`
	Sh []SendHistory       `json:"send_history"`
	Rh []ReceiveHistory    `json:"receive_history"`
}

type DepositHistory struct {
	Time   string  `json:"time"`
	Amount float32 `json:"amount"`
}
type WithdrawalHistory struct {
	Time   string  `json:"time"`
	Amount float32 `json:"amount"`
}
type SendHistory struct {
	ReceiverID int     `json:"receiver_id"`
	Time       string  `json:"time"`
	Amount     float32 `json:"amount"`
}
type ReceiveHistory struct {
	SenderID int     `json:"sender_id"`
	Time     string  `json:"time"`
	Amount   float32 `json:"amount"`
}

func (t *TransactionHistory) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t)
	return j, err
}
