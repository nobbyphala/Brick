// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/usecase.go
//
// Generated by this command:
//
//	mockgen -source=./usecase/usecase.go -destination=./mock/usecase/usecase.go -package=mock_usecase
//

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	domain "github.com/nobbyphala/Brick/domain"
	usecase "github.com/nobbyphala/Brick/usecase"
	gomock "go.uber.org/mock/gomock"
)

// MockDisbursement is a mock of Disbursement interface.
type MockDisbursement struct {
	ctrl     *gomock.Controller
	recorder *MockDisbursementMockRecorder
}

// MockDisbursementMockRecorder is the mock recorder for MockDisbursement.
type MockDisbursementMockRecorder struct {
	mock *MockDisbursement
}

// NewMockDisbursement creates a new mock instance.
func NewMockDisbursement(ctrl *gomock.Controller) *MockDisbursement {
	mock := &MockDisbursement{ctrl: ctrl}
	mock.recorder = &MockDisbursementMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDisbursement) EXPECT() *MockDisbursementMockRecorder {
	return m.recorder
}

// Disburse mocks base method.
func (m *MockDisbursement) Disburse(ctx context.Context, disbursement domain.Disbursement) (domain.Disbursement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disburse", ctx, disbursement)
	ret0, _ := ret[0].(domain.Disbursement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Disburse indicates an expected call of Disburse.
func (mr *MockDisbursementMockRecorder) Disburse(ctx, disbursement any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disburse", reflect.TypeOf((*MockDisbursement)(nil).Disburse), ctx, disbursement)
}

// ProcessBankCallback mocks base method.
func (m *MockDisbursement) ProcessBankCallback(ctx context.Context, bankCallback usecase.BankCallbackData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessBankCallback", ctx, bankCallback)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessBankCallback indicates an expected call of ProcessBankCallback.
func (mr *MockDisbursementMockRecorder) ProcessBankCallback(ctx, bankCallback any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBankCallback", reflect.TypeOf((*MockDisbursement)(nil).ProcessBankCallback), ctx, bankCallback)
}

// VerifyDisbursement mocks base method.
func (m *MockDisbursement) VerifyDisbursement(ctx context.Context, disbursement domain.Disbursement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyDisbursement", ctx, disbursement)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyDisbursement indicates an expected call of VerifyDisbursement.
func (mr *MockDisbursementMockRecorder) VerifyDisbursement(ctx, disbursement any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyDisbursement", reflect.TypeOf((*MockDisbursement)(nil).VerifyDisbursement), ctx, disbursement)
}
