package api

import "encoding/json"

type UserBalance struct {
	ID      int     `json:"id"`
	Balance float32 `json:"balance"`
}

func (u *UserBalance) GetJSON() ([]byte, error) {
	j, err := json.Marshal(u)
	return j, err
}

type Transaction struct {
	ID       int     `json:"id"`
	Receiver int     `json:"receiver"`
	Amount   float32 `json:"amount"`
}
