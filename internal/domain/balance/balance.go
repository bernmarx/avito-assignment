package balance

import "encoding/json"

type Balance struct {
	Storage         StorageAccess
	ExchangeRateApi ExchangeRateGetter
}

func NewBalance(s StorageAccess, eR ExchangeRateGetter) *Balance {
	return &Balance{s, eR}
}

func (b *Balance) MakeDeposit(account_id int, balance_id int, amount float32) error {
	err := b.Storage.DepositMoney(account_id, balance_id, amount)

	return err
}
func (b *Balance) MakeWithdraw(account_id int, balance_id int, amount float32) error {
	err := b.Storage.WithdrawMoney(account_id, balance_id, amount)

	return err
}
func (b *Balance) MakeTransfer(sender_account_id int, sender_balance_id int,
	receiver_account_id int, receiver_balance_id int, amount float32) error {

	err := b.Storage.TransferMoney(sender_account_id, sender_balance_id, receiver_account_id,
		receiver_balance_id, amount)

	return err
}
func (b *Balance) GetBalance(account_id int, balance_id int) (float32, error) {
	bal, err := b.Storage.GetBalance(account_id, balance_id)
	return bal, err
}

func (b *Balance) GetBalanceHistory(account_id int, balance_id int, sort string, page int64) ([]byte, error) {
	h, err := b.Storage.GetBalanceHistory(account_id, balance_id, sort, page)
	if err != nil {
		return make([]byte, 0), err
	}

	j, err := json.Marshal(h)

	return j, err
}
