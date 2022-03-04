//go:generate mockgen -source $GOFILE -destination ./balance_mock.go -package $GOPACKAGE
package balance

import "github.com/bernmarx/avito-assignment/internal/storage"

type storageIfc interface {
	Deposit(id int, amount float32) error
	Withdraw(id int, amount float32) error
	Transfer(senderID int, receiverID int, amount float32) error
	GetBalance(id int) (float32, error)
	GetTransactionHistory(id int) (storage.TransactionHistory, error)
	GetTransactionHistoryPage(id int, sort string, page int) (storage.TransactionHistory, error)
}

type Balance struct {
	Storage storageIfc
}

func NewBalance() *Balance {
	s, _ := storage.NewStorage()
	return &Balance{s}
}
func (b *Balance) SetInterface(ifc storageIfc) {
	b.Storage = ifc
}
