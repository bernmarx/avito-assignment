package balance

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMockstorageIfc(ctrl)

	m.EXPECT().Deposit(int(123), float32(10.0)).Return(nil)

	b := Balance{m}
	err := b.Deposit(123, 10.0)

	assert.Nil(t, err)

	m.EXPECT().Deposit(int(100), float32(1.0)).Return(errors.New("some error"))

	err = b.Deposit(100, 1.0)

	assert.NotNil(t, err)

	err = b.Deposit(0, 3.0)

	assert.NotNil(t, err)

	err = b.Deposit(11, 0.0)

	assert.NotNil(t, err)
}

func TestWithdraw(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMockstorageIfc(ctrl)

	m.EXPECT().Withdraw(int(123), float32(10.0)).Return(nil)

	b := Balance{m}
	err := b.Withdraw(123, 10.0)

	assert.Nil(t, err)

	m.EXPECT().Withdraw(int(100), float32(1.0)).Return(errors.New("some error"))

	err = b.Withdraw(100, 1.0)

	assert.NotNil(t, err)

	err = b.Withdraw(0, 3.0)

	assert.NotNil(t, err)

	err = b.Withdraw(11, 0.0)

	assert.NotNil(t, err)
}

func TestTransfer(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMockstorageIfc(ctrl)

	m.EXPECT().Transfer(int(12), int(34), float32(100.0)).Return(nil)

	b := Balance{m}
	err := b.Transfer(12, 34, 100.0)

	assert.Nil(t, err)

	m.EXPECT().Transfer(int(56), int(78), float32(1.0)).Return(errors.New("some error"))

	err = b.Transfer(56, 78, 1.0)

	assert.NotNil(t, err)

	err = b.Transfer(0, 10, 24.5)

	assert.NotNil(t, err)

	err = b.Transfer(10, 0, 24.5)

	assert.NotNil(t, err)

	err = b.Transfer(10, 20, 0)

	assert.NotNil(t, err)
}

func TestGetBalance(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMockstorageIfc(ctrl)

	m.EXPECT().GetBalance(int(10)).Return(float32(24.0), nil)

	b := Balance{m}

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
