// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/storage (interfaces: AccountsClient)

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	storage "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-09-01/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountsClient is a mock of AccountsClient interface.
type MockAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsClientMockRecorder
}

// MockAccountsClientMockRecorder is the mock recorder for MockAccountsClient.
type MockAccountsClientMockRecorder struct {
	mock *MockAccountsClient
}

// NewMockAccountsClient creates a new mock instance.
func NewMockAccountsClient(ctrl *gomock.Controller) *MockAccountsClient {
	mock := &MockAccountsClient{ctrl: ctrl}
	mock.recorder = &MockAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountsClient) EXPECT() *MockAccountsClientMockRecorder {
	return m.recorder
}

// GetProperties mocks base method.
func (m *MockAccountsClient) GetProperties(arg0 context.Context, arg1, arg2 string, arg3 storage.AccountExpand) (storage.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProperties", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(storage.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProperties indicates an expected call of GetProperties.
func (mr *MockAccountsClientMockRecorder) GetProperties(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProperties", reflect.TypeOf((*MockAccountsClient)(nil).GetProperties), arg0, arg1, arg2, arg3)
}

// ListAccountSAS mocks base method.
func (m *MockAccountsClient) ListAccountSAS(arg0 context.Context, arg1, arg2 string, arg3 storage.AccountSasParameters) (storage.ListAccountSasResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccountSAS", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(storage.ListAccountSasResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountSAS indicates an expected call of ListAccountSAS.
func (mr *MockAccountsClientMockRecorder) ListAccountSAS(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountSAS", reflect.TypeOf((*MockAccountsClient)(nil).ListAccountSAS), arg0, arg1, arg2, arg3)
}

// ListKeys mocks base method.
func (m *MockAccountsClient) ListKeys(arg0 context.Context, arg1, arg2 string, arg3 storage.ListKeyExpand) (storage.AccountListKeysResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(storage.AccountListKeysResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockAccountsClientMockRecorder) ListKeys(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockAccountsClient)(nil).ListKeys), arg0, arg1, arg2, arg3)
}

// Update mocks base method.
func (m *MockAccountsClient) Update(arg0 context.Context, arg1, arg2 string, arg3 storage.AccountUpdateParameters) (storage.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(storage.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAccountsClientMockRecorder) Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccountsClient)(nil).Update), arg0, arg1, arg2, arg3)
}
