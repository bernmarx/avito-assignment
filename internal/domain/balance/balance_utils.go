package balance

import (
	"errors"
	"log"
)

func checkID(id int) error {
	if id <= 0 {
		log.Println("found invalid ID")
		return errors.New("missing or invalid ID")
	}

	return nil
}
func checkAmount(amount float32) error {
	if amount <= 0 {
		log.Println("found invalid amount")
		return errors.New("invalid amount")
	}

	return nil
}
