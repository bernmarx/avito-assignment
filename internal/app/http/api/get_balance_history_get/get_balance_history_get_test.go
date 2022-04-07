package get_balance_history_get

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/bernmarx/avito-assignment/internal/domain/balance"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func TestHandler(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	s := balance.NewMockStorageAccess(ctrl)

	handler := http.HandlerFunc(Handler(s))

	s.EXPECT().GetBalanceHistory(10, 20, "value_asc", int64(0)).Return([]byte(`test`), nil)

	rw := httptest.NewRecorder()
	body := strings.NewReader(`{"account_id": 10, "balance_id": 20, "sort":"value_asc"}`)
	req, err := http.NewRequest("GET", "/history", body)

	assert.Nil(t, err)

	handler.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)
	assert.Equal(t, rw.Body.String(), `test`)

	s.EXPECT().GetBalanceHistory(10, 20, "value_asc", int64(0)).Return(nil, errors.New("msg", 500))

	rw = httptest.NewRecorder()
	body = strings.NewReader(`{"account_id": 10, "balance_id": 20, "sort":"value_asc"}`)
	req, err = http.NewRequest("GET", "/history", body)

	assert.Nil(t, err)

	handler.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusInternalServerError, rw.Code)
}
