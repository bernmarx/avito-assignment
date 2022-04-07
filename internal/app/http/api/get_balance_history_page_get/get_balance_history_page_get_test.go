package get_balance_history_page_get

import (
	"testing"
)

func TestHandler(t *testing.T) {
	// ctrl := gomock.NewController(t)

	// defer ctrl.Finish()

	// s := balance.NewMockStorageAccess(ctrl)

	// ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, `
	// 	{"conversion_rate":2.0}
	// 	`)
	// }))

	// eR := exchangerateclient.NewExchangeRate(http.DefaultClient, ts.URL, 0)

	// r := mux.NewRouter()
	// r.HandleFunc("/history/{page}", Handler(s, eR)).Methods("GET")

	// testServer := httptest.NewServer(r)

	// s.EXPECT().GetBalanceHistory(10, 20, "value_desc", int64(2)).Return([]byte(`test`), nil)

	// rw := httptest.NewRecorder()
	// body := strings.NewReader(`{"account_id": 10, "balance_id": 20, "sort":"value_desc"}`)
	// req, err := http.NewRequest("GET", "/history/2", body)

	// assert.Nil(t, err)

	// handler.ServeHTTP(rw, req)

	// assert.Equal(t, http.StatusOK, rw.Code)
	// assert.Equal(t, rw.Body.String(), `test`)

	// s.EXPECT().GetBalanceHistory(10, 20, "value_asc", int64(1)).Return(nil, errors.New("msg", 500))

	// rw = httptest.NewRecorder()
	// body = strings.NewReader(`{"account_id": 10, "balance_id": 20, "sort":"value_asc"}`)
	// req, err = http.NewRequest("GET", "/history/1", body)

	// assert.Nil(t, err)

	// handler.ServeHTTP(rw, req)

	// assert.Equal(t, http.StatusInternalServerError, rw.Code)
}
