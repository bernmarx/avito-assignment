package balance

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMakeDeposit(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMockStorageAccess(ctrl)

	meR := NewMockExchangeRateGetter(ctrl)

	m.EXPECT().DepositMoney(int(123), float32(10.0)).Return(nil)

	b := Balance{m, meR}
	err := b.MakeDeposit(123, 10.0)

	assert.Nil(t, err)

	m.EXPECT().DepositMoney(int(100), float32(1.0)).Return(errors.New("some error"))

	err = b.MakeDeposit(100, 1.0)

	assert.NotNil(t, err)

	err = b.MakeDeposit(0, 3.0)

	assert.NotNil(t, err)

	err = b.MakeDeposit(11, 0.0)

	assert.NotNil(t, err)
}

func TestMakeWithdraw(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMockStorageAccess(ctrl)

	meR := NewMockExchangeRateGetter(ctrl)

	m.EXPECT().WithdrawMoney(int(123), float32(10.0)).Return(nil)

	b := Balance{m, meR}
	err := b.MakeWithdraw(123, 10.0)

	assert.Nil(t, err)

	m.EXPECT().WithdrawMoney(int(100), float32(1.0)).Return(errors.New("some error"))

	err = b.MakeWithdraw(100, 1.0)

	assert.NotNil(t, err)

	err = b.MakeWithdraw(0, 3.0)

	assert.NotNil(t, err)

	err = b.MakeWithdraw(11, 0.0)

	assert.NotNil(t, err)
}

func TestMakeTransfer(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMockStorageAccess(ctrl)

	meR := NewMockExchangeRateGetter(ctrl)

	m.EXPECT().TransferMoney(int(12), int(34), float32(100.0)).Return(nil)

	b := Balance{m, meR}
	err := b.MakeTransfer(12, 34, 100.0)

	assert.Nil(t, err)

	m.EXPECT().TransferMoney(int(56), int(78), float32(1.0)).Return(errors.New("some error"))

	err = b.MakeTransfer(56, 78, 1.0)

	assert.NotNil(t, err)

	err = b.MakeTransfer(0, 10, 24.5)

	assert.NotNil(t, err)

	err = b.MakeTransfer(10, 0, 24.5)

	assert.NotNil(t, err)

	err = b.MakeTransfer(10, 20, 0)

	assert.NotNil(t, err)
}

func TestGetBalance(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMockStorageAccess(ctrl)

	meR := NewMockExchangeRateGetter(ctrl)

	m.EXPECT().GetBalance(int(10)).Return(float32(24.0), nil)

	b := Balance{m, meR}

	bal, err := b.GetBalance(10)

	assert.Equal(t, float32(24.0), bal)
	assert.Nil(t, err)

	m.EXPECT().GetBalance(int(30)).Return(float32(0.0), errors.New("some error"))

	bal, err = b.GetBalance(30)

	assert.Equal(t, float32(0.0), bal)
	assert.NotNil(t, err)

	bal, err = b.GetBalance(0)

	assert.Equal(t, float32(0.0), bal)
	assert.NotNil(t, err)
}
