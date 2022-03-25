//go:generate mockgen -source $GOFILE -destination ./storage_mock.go -package $GOPACKAGE
package balance

import (
	"database/sql"
	"encoding/json"
)

type Transaction struct {
	Operation         string        `json:"operation"`
	CreatedAt         string        `json:"created_at"`
	Value             float32       `json:"value"`
	ReceiverAccountID sql.NullInt64 `json:"receiver_account_id"`
	SenderAccountID   sql.NullInt64 `json:"sender_account_id"`
}

func (t *Transaction) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t)
	return j, err
}
