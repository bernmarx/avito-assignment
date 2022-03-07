// Code generated by MockGen. DO NOT EDIT.
// Source: converter.go

// Package converter is a generated GoMock package.
package converter

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockconversionRate is a mock of conversionRate interface
type MockconversionRate struct {
	ctrl     *gomock.Controller
	recorder *MockconversionRateMockRecorder
}

// MockconversionRateMockRecorder is the mock recorder for MockconversionRate
type MockconversionRateMockRecorder struct {
	mock *MockconversionRate
}

// NewMockconversionRate creates a new mock instance
func NewMockconversionRate(ctrl *gomock.Controller) *MockconversionRate {
	mock := &MockconversionRate{ctrl: ctrl}
	mock.recorder = &MockconversionRateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockconversionRate) EXPECT() *MockconversionRateMockRecorder {
	return m.recorder
}

// GetExchangeRate mocks base method
func (m *MockconversionRate) GetExchangeRate(baseCur, cur string) (float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchangeRate", baseCur, cur)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate
func (mr *MockconversionRateMockRecorder) GetExchangeRate(baseCur, cur interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockconversionRate)(nil).GetExchangeRate), baseCur, cur)
}