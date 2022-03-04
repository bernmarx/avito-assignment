package balance

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
