package exchangerateapi

import (
	"encoding/json"
	"net/http"
	"os"
)

type ExchangeRate struct {
	httpClient
}

func NewExchangeRate() *ExchangeRate {
	return &ExchangeRate{httpClient: http.DefaultClient}
}

func (e *ExchangeRate) GetExchangeRate(baseCur string, cur string) (float32, error) {
	exchangeAPIkey := os.Getenv("ER_API_KEY")
	resp, err := e.Get("https://v6.exchangerate-api.com/v6/" + exchangeAPIkey + "/pair/" +
		baseCur + "/" + cur)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	type conversion struct {
		Rate float32 `json:"conversion_rate"`
	}

	var cr conversion

	err = json.NewDecoder(resp.Body).Decode(&cr)

	return cr.Rate, err
}
