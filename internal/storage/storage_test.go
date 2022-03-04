package storage

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockstorage := NewMockstorage(ctrl)
	mockstoragerow := NewMockstorageRow(ctrl)
	mockstoragerows := NewMockstorageRows(ctrl)

	s := Storage{mockstorage, mockstoragerow, mockstoragerows}

	mockstorage.EXPECT().Exec(`call balance_deposit($1, $2)`, int(10), float32(33.0)).Return(nil, nil)

	err := s.Deposit(10, 33.0)

	assert.Nil(t, err)

	mockstorage.EXPECT().Exec(`call balance_deposit($1, $2)`, int(22), float32(33.0)).Return(nil, errors.New("err"))

	err = s.Deposit(22, 33.0)

	assert.NotNil(t, err)
}

func TestWithdraw(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockstorage := NewMockstorage(ctrl)
	mockstoragerow := NewMockstorageRow(ctrl)
	mockstoragerows := NewMockstorageRows(ctrl)

	s := Storage{mockstorage, mockstoragerow, mockstoragerows}

	mockstorage.EXPECT().Exec(`call balance_withdraw($1, $2)`, int(10), float32(33.0)).Return(nil, nil)

	err := s.Withdraw(10, 33.0)

	assert.Nil(t, err)

	mockstorage.EXPECT().Exec(`call balance_withdraw($1, $2)`, int(22), float32(33.0)).Return(nil, errors.New("err"))

	err = s.Withdraw(22, 33.0)

	assert.NotNil(t, err)
}

func TestTransfer(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockstorage := NewMockstorage(ctrl)
	mockstoragerow := NewMockstorageRow(ctrl)
	mockstoragerows := NewMockstorageRows(ctrl)

	s := Storage{mockstorage, mockstoragerow, mockstoragerows}

	mockstorage.EXPECT().Exec(`call balance_transfer($1, $2, $3)`, int(10), int(20), float32(33.0)).Return(nil, nil)

	err := s.Transfer(10, 20, 33.0)

	assert.Nil(t, err)

	mockstorage.EXPECT().Exec(`call balance_transfer($1, $2, $3)`, int(22), int(33), float32(33.0)).Return(nil, errors.New("err"))

	err = s.Transfer(22, 33, 33.0)

	assert.NotNil(t, err)
}
