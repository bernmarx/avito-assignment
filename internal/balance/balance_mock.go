// Code generated by MockGen. DO NOT EDIT.
// Source: balance_model.go

// Package balance is a generated GoMock package.
package balance

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorageAccess is a mock of StorageAccess interface
type MockStorageAccess struct {
	ctrl     *gomock.Controller
	recorder *MockStorageAccessMockRecorder
}

// MockStorageAccessMockRecorder is the mock recorder for MockStorageAccess
type MockStorageAccessMockRecorder struct {
	mock *MockStorageAccess
}

// NewMockStorageAccess creates a new mock instance
func NewMockStorageAccess(ctrl *gomock.Controller) *MockStorageAccess {
	mock := &MockStorageAccess{ctrl: ctrl}
	mock.recorder = &MockStorageAccessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageAccess) EXPECT() *MockStorageAccessMockRecorder {
	return m.recorder
}

// DepositMoney mocks base method
func (m *MockStorageAccess) DepositMoney(id int, amount float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DepositMoney", id, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// DepositMoney indicates an expected call of DepositMoney
func (mr *MockStorageAccessMockRecorder) DepositMoney(id, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepositMoney", reflect.TypeOf((*MockStorageAccess)(nil).DepositMoney), id, amount)
}

// WithdrawMoney mocks base method
func (m *MockStorageAccess) WithdrawMoney(id int, amount float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawMoney", id, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithdrawMoney indicates an expected call of WithdrawMoney
func (mr *MockStorageAccessMockRecorder) WithdrawMoney(id, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawMoney", reflect.TypeOf((*MockStorageAccess)(nil).WithdrawMoney), id, amount)
}

// TransferMoney mocks base method
func (m *MockStorageAccess) TransferMoney(senderID, receiverID int, amount float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferMoney", senderID, receiverID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferMoney indicates an expected call of TransferMoney
func (mr *MockStorageAccessMockRecorder) TransferMoney(senderID, receiverID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferMoney", reflect.TypeOf((*MockStorageAccess)(nil).TransferMoney), senderID, receiverID, amount)
}

// GetBalance mocks base method
func (m *MockStorageAccess) GetBalance(id int) (float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", id)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance
func (mr *MockStorageAccessMockRecorder) GetBalance(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockStorageAccess)(nil).GetBalance), id)
}

// GetTransactionHistory mocks base method
func (m *MockStorageAccess) GetTransactionHistory(id int) (TransactionHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionHistory", id)
	ret0, _ := ret[0].(TransactionHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionHistory indicates an expected call of GetTransactionHistory
func (mr *MockStorageAccessMockRecorder) GetTransactionHistory(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionHistory", reflect.TypeOf((*MockStorageAccess)(nil).GetTransactionHistory), id)
}

// GetTransactionHistoryPage mocks base method
func (m *MockStorageAccess) GetTransactionHistoryPage(id int, sort string, page int) (TransactionHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionHistoryPage", id, sort, page)
	ret0, _ := ret[0].(TransactionHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionHistoryPage indicates an expected call of GetTransactionHistoryPage
func (mr *MockStorageAccessMockRecorder) GetTransactionHistoryPage(id, sort, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionHistoryPage", reflect.TypeOf((*MockStorageAccess)(nil).GetTransactionHistoryPage), id, sort, page)
}
