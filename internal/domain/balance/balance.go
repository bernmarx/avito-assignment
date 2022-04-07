package balance

// Balance represents all available manipulations with users' balance
type Balance struct {
	Storage StorageAccess
}

// NewBalance creates new Balance
func NewBalance(s StorageAccess) *Balance {
	return &Balance{s}
}
