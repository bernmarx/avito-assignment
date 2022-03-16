//go:generate mockgen -source $GOFILE -destination ./exchangerate_mock.go -package $GOPACKAGE
package exchangerateclient

import (
	"encoding/json"
	"net/http"

	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

type HttpClient interface {
	Get(url string) (*http.Response, error)
}

type ExchangeRate struct {
	HttpClient
	url    string
	curPos int
}

type rate struct {
	Value float32 `json:"conversion_rate"`
}

func NewExchangeRate(c HttpClient, _url string, _curPos int) *ExchangeRate {
	return &ExchangeRate{HttpClient: c, url: _url, curPos: _curPos}
}

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
