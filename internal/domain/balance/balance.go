package balance

type Balance struct {
	Storage         StorageAccess
	ExchangeRateApi ExchangeRateGetter
}

func NewBalance(s StorageAccess, eR ExchangeRateGetter) *Balance {
	return &Balance{s, eR}
}
