//go:generate mockgen -source $GOFILE -destination ./exchangerate_mock.go -package $GOPACKAGE

package exchangerateclient

import (
	"encoding/json"
	"net/http"

	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

// HTTPClientGetter provides access to making a request to 'url' with GET method
type HTTPClientGetter interface {
	Get(url string) (*http.Response, error)
}

// ExchangeRate is an entity that manages getting exchange rate for currencies
type ExchangeRate struct {
	HTTPClientGetter
	url    string
	curPos int
}

type rate struct {
	Value float32 `json:"conversion_rate"`
}

// NewExchangeRate creates new ExchangeRate
func NewExchangeRate(c HTTPClientGetter, _url string, _curPos int) *ExchangeRate {
	return &ExchangeRate{HTTPClientGetter: c, url: _url, curPos: _curPos}
}

// GetExchangeRate returns exchange rate for 'cur' currency
func (e *ExchangeRate) GetExchangeRate(cur string) (float32, error) {
	url := e.url[:e.curPos] + cur + e.url[e.curPos:]
	resp, err := e.Get(url)
	if err != nil {
		return 0.0, errors.New(err.Error(), 500)
	}
	defer resp.Body.Close()

	r := rate{}

	err = json.NewDecoder(resp.Body).Decode(&r)

	if err != nil {
		return 0.0, errors.New(err.Error(), 500)
	}

	return r.Value, nil
}
