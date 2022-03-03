package balance

func (b *Balance) GetBalance(id int) (float32, error) {
	err := checkID(id)
	if err != nil {
		return 0.0, err
	}
	bal, err := b.Storage.GetBalance(id)
	return bal, err
}
