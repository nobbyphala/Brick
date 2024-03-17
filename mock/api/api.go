// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/api/api.go
//
// Generated by this command:
//
//	mockgen -source=./usecase/api/api.go -destination=./mock/api/api.go -package=mock_api
//

// Package mock_api is a generated GoMock package.
package mock_api

import (
	context "context"
	reflect "reflect"

	api "github.com/nobbyphala/Brick/usecase/api"
	gomock "go.uber.org/mock/gomock"
)

// MockBank is a mock of Bank interface.
type MockBank struct {
	ctrl     *gomock.Controller
	recorder *MockBankMockRecorder
}

// MockBankMockRecorder is the mock recorder for MockBank.
type MockBankMockRecorder struct {
	mock *MockBank
}

// NewMockBank creates a new mock instance.
func NewMockBank(ctrl *gomock.Controller) *MockBank {
	mock := &MockBank{ctrl: ctrl}
	mock.recorder = &MockBankMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBank) EXPECT() *MockBankMockRecorder {
	return m.recorder
}

// TransferMoney mocks base method.
func (m *MockBank) TransferMoney(ctx context.Context, transfer api.TransferRequest) (api.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferMoney", ctx, transfer)
	ret0, _ := ret[0].(api.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferMoney indicates an expected call of TransferMoney.
func (mr *MockBankMockRecorder) TransferMoney(ctx, transfer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferMoney", reflect.TypeOf((*MockBank)(nil).TransferMoney), ctx, transfer)
}

// VerifyAccount mocks base method.
func (m *MockBank) VerifyAccount(ctx context.Context, account api.VerifyAccountRequest) (api.VerifyAccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyAccount", ctx, account)
	ret0, _ := ret[0].(api.VerifyAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyAccount indicates an expected call of VerifyAccount.
func (mr *MockBankMockRecorder) VerifyAccount(ctx, account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAccount", reflect.TypeOf((*MockBank)(nil).VerifyAccount), ctx, account)
}