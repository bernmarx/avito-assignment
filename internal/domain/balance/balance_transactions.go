package balance

func (b *Balance) Deposit(id int, amount float32) error {
	err := checkID(id)
	if err != nil {
		return err
	}
	err = checkAmount(amount)
	if err != nil {
		return err
	}

	err = b.Storage.Deposit(id, amount)

	return err
}
func (b *Balance) Withdraw(id int, amount float32) error {
	err := checkID(id)
	if err != nil {
		return err
	}
	err = checkAmount(amount)
	if err != nil {
		return err
	}

	err = b.Storage.Withdraw(id, amount)

	return err
}
func (b *Balance) Transfer(id int, receiverId int, amount float32) error {
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

	err = b.Storage.Transfer(id, receiverId, amount)

	return err
}
