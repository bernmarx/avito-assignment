package api

import "encoding/json"

type RequestData struct {
	ID      int     `json:"id"`
	Balance float32 `json:"balance"`
	Sort    string  `json:"sort"`
}

func (u *RequestData) GetJSON() ([]byte, error) {
	j, err := json.Marshal(u)
	return j, err
}

type Transaction struct {
	ID       int     `json:"id"`
	Receiver int     `json:"receiver"`
	Amount   float32 `json:"amount"`
}
