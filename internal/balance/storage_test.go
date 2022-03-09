package balance

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

//Rewrite tests they do not match methods' logic

func TestWithdraw(t *testing.T) {
	db, mock, err := sqlmock.New()

	assert.Nil(t, err)

	s := NewStorage(db)

	mock.ExpectExec("call balance_withdraw*").
		WithArgs(int(10), float32(33.0)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = s.WithdrawMoney(10, 33.0)

	assert.Nil(t, err)

	mock.ExpectExec("call balance_withdraw*").
		WithArgs(int(3), float32(1.0)).
		WillReturnError(errors.New("some error"))

	err = s.WithdrawMoney(3, 1.0)

	assert.NotNil(t, err)
}

func TestTransfer(t *testing.T) {
	db, mock, err := sqlmock.New()

	assert.Nil(t, err)

	s := NewStorage(db)

	mock.ExpectExec("call balance_transfer*").
		WithArgs(int(1), int(2), float32(3.0)).
		WillReturnResult(sqlmock.NewResult(0, 2))

	err = s.TransferMoney(1, 2, 3.0)

	assert.Nil(t, err)

	mock.ExpectExec("call balance_transfer*").
		WithArgs(int(24), int(22), float32(35.0)).
		WillReturnError(errors.New("some error"))

	err = s.TransferMoney(24, 22, 35.0)

	assert.NotNil(t, err)
}
