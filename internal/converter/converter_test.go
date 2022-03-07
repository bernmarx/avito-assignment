package converter

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConvertCurrency(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMockconversionRate(ctrl)

	m.EXPECT().GetExchangeRate("RUB", "USD").Return(float32(0.1), nil)

	os.Setenv("BASE_CURRENCY", "RUB")

	c := NewConverter(m)

	got, err := c.ConvertCurrency(10.0, "USD")

	assert.Nil(t, err)
	assert.Equal(t, float32(1), got)

	m.EXPECT().GetExchangeRate("RUB", "EUR").Return(float32(0.0), errors.New("some error"))

	c = NewConverter(m)

	got, err = c.ConvertCurrency(123.4, "EUR")

	assert.NotNil(t, err)
	assert.Equal(t, float32(0.0), got)

	got, err = c.ConvertCurrency(0.0, "USD")

	assert.NotNil(t, err)
	assert.Equal(t, float32(0.0), got)

	got, err = c.ConvertCurrency(100.0, "ABCDE")

	assert.NotNil(t, err)
	assert.Equal(t, float32(0.0), got)
}
