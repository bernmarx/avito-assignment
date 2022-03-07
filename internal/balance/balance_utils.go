package balance

import (
	"errors"
)

func checkID(id int) error {
	if id <= 0 {
		return errors.New("missing or invalid ID")
	}

	return nil
}
func checkAmount(amount float32) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}

	return nil
}
