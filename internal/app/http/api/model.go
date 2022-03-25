package api

type DepositWithdrawRequestData struct {
	Account_id int     `json:"account_id"`
	Balance_id int     `json:"balance_id"`
	Amount     float32 `json:"amount"`
}

type TransferRequestData struct {
	Sender_account_id   int     `json:"sender_account_id"`
	Sender_balance_id   int     `json:"sender_balance_id"`
	Receiver_account_id int     `json:"receiver_account_id"`
	Receiver_balance_id int     `json:"receiver_balance_id"`
	Amount              float32 `json:"amount"`
}

type GetBalanceRequestData struct {
	Account_id int `json:"account_id"`
	Balance_id int `json:"balance_id"`
}

type GetBalanceHistoryRequestData struct {
	Account_id int    `json:"account_id"`
	Balance_id int    `json:"balance_id"`
	Sort       string `json:"sort"`
}

type Balance struct {
	Account_id int     `json:"account_id"`
	Balance_id int     `json:"balance_id"`
	Balance    float32 `json:"balance"`
}
