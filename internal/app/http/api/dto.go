package api

// DepositWithdrawRequestData represents request data for deposit and withdraw handlers
type DepositWithdrawRequestData struct {
	AccountID int     `json:"account_id"`
	BalanceID int     `json:"balance_id"`
	Amount    float32 `json:"amount"`
}

// TransferRequestData represents request data for transfer handler
type TransferRequestData struct {
	SenderAccountID   int     `json:"sender_account_id"`
	SenderBalanceID   int     `json:"sender_balance_id"`
	ReceiverAccountID int     `json:"receiver_account_id"`
	ReceiverBalanceID int     `json:"receiver_balance_id"`
	Amount            float32 `json:"amount"`
}

// GetBalanceRequestData represents request data for get_balance handler
type GetBalanceRequestData struct {
	AccountID int `json:"account_id"`
	BalanceID int `json:"balance_id"`
}

// GetBalanceHistoryRequestData represents request data for get_balance_history handler
type GetBalanceHistoryRequestData struct {
	AccountID int    `json:"account_id"`
	BalanceID int    `json:"balance_id"`
	Sort      string `json:"sort"`
}

// Balance represents request data for get_balance handler
type Balance struct {
	AccountID int     `json:"account_id"`
	BalanceID int     `json:"balance_id"`
	Balance   float32 `json:"balance"`
}
