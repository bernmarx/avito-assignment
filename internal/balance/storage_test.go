package balance

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockstorage := NewMockdatabase(ctrl)

	s := Storage{mockstorage}

	mockstorage.EXPECT().Exec(`call balance_deposit($1, $2)`, int(10), float32(33.0)).Return(nil, nil)

	err := s.DepositMoney(10, 33.0)

	assert.Nil(t, err)

	mockstorage.EXPECT().Exec(`call balance_deposit($1, $2)`, int(22), float32(33.0)).Return(nil, errors.New("err"))

	err = s.DepositMoney(22, 33.0)

	assert.NotNil(t, err)
}

func TestWithdraw(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockstorage := NewMockdatabase(ctrl)

	s := Storage{mockstorage}

	mockstorage.EXPECT().Exec(`call balance_withdraw($1, $2)`, int(10), float32(33.0)).Return(nil, nil)

	err := s.WithdrawMoney(10, 33.0)

	assert.Nil(t, err)

	mockstorage.EXPECT().Exec(`call balance_withdraw($1, $2)`, int(22), float32(33.0)).Return(nil, errors.New("err"))

	err = s.WithdrawMoney(22, 33.0)

	assert.NotNil(t, err)
}

func TestTransfer(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockstorage := NewMockdatabase(ctrl)

	s := Storage{mockstorage}

	mockstorage.EXPECT().Exec(`call balance_transfer($1, $2, $3)`, int(10), int(20), float32(33.0)).Return(nil, nil)

	err := s.TransferMoney(10, 20, 33.0)

	assert.Nil(t, err)

	mockstorage.EXPECT().Exec(`call balance_transfer($1, $2, $3)`, int(22), int(33), float32(33.0)).Return(nil, errors.New("err"))

	err = s.TransferMoney(22, 33, 33.0)

	assert.NotNil(t, err)
}
