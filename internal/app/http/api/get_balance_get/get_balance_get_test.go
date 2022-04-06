package get_balance_get

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bernmarx/avito-assignment/internal/app/http/api"
	"github.com/bernmarx/avito-assignment/internal/domain/balance"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/exchangerateclient"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	s := balance.NewMockStorageAccess(ctrl)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		{"conversion_rate":2.0}
		`)
	}))

	eR := exchangerateclient.NewExchangeRate(http.DefaultClient, ts.URL, 0)

	handler := http.HandlerFunc(Handler(s, eR))

	s.EXPECT().GetBalance(10, 20).Return(float32(15.0), nil)

	rw := httptest.NewRecorder()
	body := strings.NewReader(`{"account_id": 10, "balance_id": 20}`)
	req, err := http.NewRequest("GET", "/balance", body)

	assert.Nil(t, err)

	handler.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)

	expected := api.Balance{AccountID: 10, BalanceID: 20, Balance: 15.0}
	var got api.Balance

	json.NewDecoder(rw.Result().Body).Decode(&got)

	assert.Equal(t, expected, got)

	s.EXPECT().GetBalance(10, 20).Return(float32(0.0), errors.New("some error", 500))

	rw = httptest.NewRecorder()
	body = strings.NewReader(`{"account_id": 10, "balance_id": 20}`)
	req, err = http.NewRequest("POST", "/deposit", body)

	assert.Nil(t, err)

	handler.ServeHTTP(rw, req)

	assert.Equal(t, 500, rw.Code)

	s.EXPECT().GetBalance(10, 20).Return(float32(15.0), nil)

	rw = httptest.NewRecorder()
	body = strings.NewReader(`{"account_id": 10, "balance_id": 20}`)
	req, err = http.NewRequest("POST", "/deposit", body)
	q := req.URL.Query()
	q.Add("currency", "")
	req.URL.RawQuery = q.Encode()

	assert.Nil(t, err)

	handler.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)

	expected = api.Balance{AccountID: 10, BalanceID: 20, Balance: 30.0}

	json.NewDecoder(rw.Result().Body).Decode(&got)

	assert.Equal(t, expected, got)

	s.EXPECT().GetBalance(10, 20).Return(float32(15.0), nil)

	rw = httptest.NewRecorder()
	body = strings.NewReader(`{"account_id": 10, "balance_id": 20}`)
	req, err = http.NewRequest("POST", "/deposit", body)
	q = req.URL.Query()
	q.Add("currency", "ASDSA")
	req.URL.RawQuery = q.Encode()

	assert.Nil(t, err)

	handler.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusInternalServerError, rw.Code)
}
