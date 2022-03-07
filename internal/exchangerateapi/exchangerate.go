//go:generate mockgen -source $GOFILE -destination ./exchangerate_mock.go -package $GOPACKAGE
package exchangerateapi

import (
	"encoding/json"
	"net/http"
)

type httpClient interface {
	Get(url string) (*http.Response, error)
}

type ExchangeRate struct {
	httpClient
	url    string
	curPos int
}

type rate struct {
	Value float32 `json:"conversion_rate"`
}

func NewExchangeRate(_url string, _curPos int) *ExchangeRate {
	return &ExchangeRate{httpClient: http.DefaultClient, url: _url, curPos: _curPos}
}

func (e *ExchangeRate) GetExchangeRate(cur string) (float32, error) {
	url := e.url[:e.curPos] + cur + e.url[e.curPos:]
	resp, err := e.Get(url)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	r := rate{}

	err = json.NewDecoder(resp.Body).Decode(&r)

	return r.Value, err
}
