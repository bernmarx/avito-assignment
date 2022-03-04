//go:generate mockgen -source $GOFILE -destination ./storage_mock.go -package $GOPACKAGE
package storage

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type storage interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type storageRow interface {
	Scan(dest ...interface{}) error
}

type storageRows interface {
	Next() bool
	Scan(dest ...interface{}) error
}

type Storage struct {
	storage
	storageRow
	storageRows
}

func NewStorage() (*Storage, error) {
	connData := "host=" + os.Getenv("DB_HOST") + " " + "port=" + os.Getenv("DB_PORT")
	connData = connData + " " + "user=" + os.Getenv("DB_USER") + " " + "password=" + os.Getenv("DB_PASSWORD")
	connData = connData + " " + "dbname=" + os.Getenv("DB_NAME") + " " + "sslmode=" + os.Getenv("DB_SSLMODE")
	log.Println(connData)

	db, err := sql.Open("postgres", connData)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	s := Storage{storage: db, storageRow: &sql.Row{}, storageRows: &sql.Rows{}}

	return &s, nil
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
