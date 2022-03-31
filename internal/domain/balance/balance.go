package balance

// Balance represents all available manipulations with users' balance
type Balance struct {
	Storage         StorageAccess
	ExchangeRateAPI ExchangeRateGetter
}

// NewBalance creates new Balance
func NewBalance(s StorageAccess, eR ExchangeRateGetter) *Balance {
	return &Balance{s, eR}
}
