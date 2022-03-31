//go:generate mockgen -source $GOFILE -destination ./balance_mock.go -package $GOPACKAGE

package balance

// StorageAccess provides access to Storage methods
type StorageAccess interface {
	DepositMoney(accountID int, balanceID int, amount float32) error
	WithdrawMoney(accountID int, balanceID int, amount float32) error
	TransferMoney(senderAccountID int, senderBalanceID int,
		receiverAccountID int, receiverBalanceID int, amount float32) error
	GetBalance(accountID int, balanceID int) (float32, error)
	GetBalanceHistory(accountID int, balanceID int, sort string, page int64) ([]byte, error)
}

// ExchangeRateGetter provides access ExchangeRate methods
type ExchangeRateGetter interface {
	GetExchangeRate(cur string) (float32, error)
}
