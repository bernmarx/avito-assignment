package balance

type Balance struct {
	Storage         StorageAccess
	ExchangeRateApi ExchangeRateGetter
}

func NewBalance(s StorageAccess, eR ExchangeRateGetter) *Balance {
	return &Balance{s, eR}
}

func (b *Balance) MakeDeposit(id int, amount float32) error {
	err := checkID(id)
	if err != nil {
		return err
	}
	err = checkAmount(amount)
	if err != nil {
		return err
	}

	err = b.Storage.DepositMoney(id, amount)

	return err
}
func (b *Balance) MakeWithdraw(id int, amount float32) error {
	err := checkID(id)
	if err != nil {
		return err
	}
	err = checkAmount(amount)
	if err != nil {
		return err
	}

	err = b.Storage.WithdrawMoney(id, amount)

	return err
}
func (b *Balance) MakeTransfer(id int, receiverId int, amount float32) error {
	err := checkID(id)
	if err != nil {
		return err
	}
	err = checkID(receiverId)
	if err != nil {
		return err
	}
	err = checkAmount(amount)
	if err != nil {
		return err
	}

	err = b.Storage.TransferMoney(id, receiverId, amount)

	return err
}
func (b *Balance) GetBalance(id int) (float32, error) {
	err := checkID(id)
	if err != nil {
		return 0.0, err
	}
	bal, err := b.Storage.GetBalance(id)
	return bal, err
}

func (b *Balance) GetTransactionHistory(id int) ([]byte, error) {
	err := checkID(id)
	if err != nil {
		return make([]byte, 0), err
	}

	t, err := b.Storage.GetTransactionHistory(id)
	if err != nil {
		return make([]byte, 0), err
	}

	j, err := t.GetJSON()

	return j, err
}

func (b *Balance) GetTransactionHistoryPage(id int, sort string, page int) ([]byte, error) {
	err := checkID(id)
	if err != nil {
		return make([]byte, 0), err
	}

	t, err := b.Storage.GetTransactionHistoryPage(id, sort, page)
	if err != nil {
		return make([]byte, 0), err
	}

	j, err := t.GetJSON()

	return j, err
}
