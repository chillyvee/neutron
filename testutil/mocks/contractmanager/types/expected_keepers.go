// Code generated by MockGen. DO NOT EDIT.
// Source: ./../../x/contractmanager/types/expected_keepers.go

// Package mock_types is a generated GoMock package.
package mock_types

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockWasmKeeper is a mock of WasmKeeper interface.
type MockWasmKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockWasmKeeperMockRecorder
}

// MockWasmKeeperMockRecorder is the mock recorder for MockWasmKeeper.
type MockWasmKeeperMockRecorder struct {
	mock *MockWasmKeeper
}

// NewMockWasmKeeper creates a new mock instance.
func NewMockWasmKeeper(ctrl *gomock.Controller) *MockWasmKeeper {
	mock := &MockWasmKeeper{ctrl: ctrl}
	mock.recorder = &MockWasmKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmKeeper) EXPECT() *MockWasmKeeperMockRecorder {
	return m.recorder
}

// HasContractInfo mocks base method.
func (m *MockWasmKeeper) HasContractInfo(ctx types.Context, contractAddress types.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasContractInfo", ctx, contractAddress)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasContractInfo indicates an expected call of HasContractInfo.
func (mr *MockWasmKeeperMockRecorder) HasContractInfo(ctx, contractAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasContractInfo", reflect.TypeOf((*MockWasmKeeper)(nil).HasContractInfo), ctx, contractAddress)
}

// Sudo mocks base method.
func (m *MockWasmKeeper) Sudo(ctx types.Context, contractAddress types.AccAddress, msg []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sudo", ctx, contractAddress, msg)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sudo indicates an expected call of Sudo.
func (mr *MockWasmKeeperMockRecorder) Sudo(ctx, contractAddress, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sudo", reflect.TypeOf((*MockWasmKeeper)(nil).Sudo), ctx, contractAddress, msg)
}