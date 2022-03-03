//go:generate mockgen -source $GOFILE -destination ./conversionrate_mock.go -package $GOPACKAGE
package conversion

import (
	"errors"
	"os"
)

type conversionRate interface {
	GetExchangeRate(baseCur string, cur string) (float32, error)
}

type Converter struct {
	conversionRate
}

func NewConverter(e conversionRate) *Converter {
	return &Converter{conversionRate: e}
}

func (c *Converter) ConvertCurrency(money float32, curr string) (float32, error) {
	if len(curr) != 3 {
		return 0.0, errors.New("invalid currency code length")
	}
	if money <= 0 {
		return 0.0, errors.New("invalid sum")
	}
	eR, err := c.GetExchangeRate(os.Getenv("BASE_CURRENCY"), curr)
	return money * eR, err
}
