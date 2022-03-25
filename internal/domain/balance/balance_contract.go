//go:generate mockgen -source $GOFILE -destination ./balance_mock.go -package $GOPACKAGE
package balance

type StorageAccess interface {
	DepositMoney(account_id int, balance_id int, amount float32) error
	WithdrawMoney(account_id int, balance_id int, amount float32) error
	TransferMoney(sender_account_id int, sender_balance_id int,
		receiver_account_id int, receiver_balance_id int, amount float32) error
	GetBalance(account_id int, balance_id int) (float32, error)
	GetBalanceHistory(account_id int, balance_id int, sort string, page int64) ([]Transaction, error)
}

type ExchangeRateGetter interface {
	GetExchangeRate(cur string) (float32, error)
}
