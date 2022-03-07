//go:generate mockgen -source $GOFILE -destination ./balance_mock.go -package $GOPACKAGE
package balance

type StorageAccess interface {
	DepositMoney(id int, amount float32) error
	WithdrawMoney(id int, amount float32) error
	TransferMoney(senderID int, receiverID int, amount float32) error
	GetBalance(id int) (float32, error)
	GetTransactionHistory(id int) (TransactionHistory, error)
	GetTransactionHistoryPage(id int, sort string, page int) (TransactionHistory, error)
}

type ExchangeRateGetter interface {
	GetExchangeRate(cur string) (float32, error)
}

type Balance struct {
	Storage         StorageAccess
	ExchangeRateApi ExchangeRateGetter
}

func NewBalance(s StorageAccess, eR ExchangeRateGetter) *Balance {
	return &Balance{s, eR}
}
