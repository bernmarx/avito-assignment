package exchangerateapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchangeRate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, `{"conversion_rate":0.5}`)
		},
	))
	defer ts.Close()

	e := NewExchangeRate(http.DefaultClient, ts.URL, len(ts.URL))

	eR, err := e.GetExchangeRate("")

	assert.Nil(t, err)
	assert.Equal(t, float32(0.5), eR)

	e = NewExchangeRate(http.DefaultClient, "garbage", 0)

	eR, err = e.GetExchangeRate("USD")

	assert.NotNil(t, err)
	assert.Equal(t, float32(0.0), eR)
}
