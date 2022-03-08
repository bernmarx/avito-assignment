package balance

import (
	"github.com/bernmarx/avito-assignment/internal/serviceerrors"
)

func checkID(id int) error {
	if id <= 0 {
		return serviceerrors.New("invalid or missing id", 400)
	}

	return nil
}
func checkAmount(amount float32) error {
	if amount <= 0 {
		return serviceerrors.New("invalid or missing amount", 400)
	}

	return nil
}
