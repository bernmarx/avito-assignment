package deposit_post

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bernmarx/avito-assignment/internal/domain/balance"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHander(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	s := balance.NewMockStorageAccess(ctrl)

	handler := http.HandlerFunc(Handler(s))

	s.EXPECT().DepositMoney(10, 20, float32(15.0)).Return(nil)

	rw := httptest.NewRecorder()
	body := strings.NewReader(`{"account_id": 10, "balance_id": 20, "amount": 15.0}`)
	req, err := http.NewRequest("POST", "/deposit", body)

	assert.Nil(t, err)

	handler.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)

	expectedBody := `Deposit was successful`

	assert.Equal(t, expectedBody, rw.Body.String())

	s.EXPECT().DepositMoney(10, 20, float32(15.0)).Return(errors.New("some error", 500))

	rw = httptest.NewRecorder()
	body = strings.NewReader(`{"account_id": 10, "balance_id": 20, "amount": 15.0}`)
	req, err = http.NewRequest("POST", "/deposit", body)

	assert.Nil(t, err)

	handler.ServeHTTP(rw, req)

	assert.Equal(t, 500, rw.Code)
}
