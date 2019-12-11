// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jim-minter/rp/pkg/util/azureclient/network (interfaces: InterfacesClient,PublicIPAddressesClient)

// Package mock_network is a generated GoMock package.
package mock_network

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-07-01/network"
	gomock "github.com/golang/mock/gomock"
)

// MockInterfacesClient is a mock of InterfacesClient interface
type MockInterfacesClient struct {
	ctrl     *gomock.Controller
	recorder *MockInterfacesClientMockRecorder
}

// MockInterfacesClientMockRecorder is the mock recorder for MockInterfacesClient
type MockInterfacesClientMockRecorder struct {
	mock *MockInterfacesClient
}

// NewMockInterfacesClient creates a new mock instance
func NewMockInterfacesClient(ctrl *gomock.Controller) *MockInterfacesClient {
	mock := &MockInterfacesClient{ctrl: ctrl}
	mock.recorder = &MockInterfacesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterfacesClient) EXPECT() *MockInterfacesClientMockRecorder {
	return m.recorder
}

// DeleteAndWait mocks base method
func (m *MockInterfacesClient) DeleteAndWait(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAndWait", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAndWait indicates an expected call of DeleteAndWait
func (mr *MockInterfacesClientMockRecorder) DeleteAndWait(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAndWait", reflect.TypeOf((*MockInterfacesClient)(nil).DeleteAndWait), arg0, arg1, arg2)
}

// MockPublicIPAddressesClient is a mock of PublicIPAddressesClient interface
type MockPublicIPAddressesClient struct {
	ctrl     *gomock.Controller
	recorder *MockPublicIPAddressesClientMockRecorder
}

// MockPublicIPAddressesClientMockRecorder is the mock recorder for MockPublicIPAddressesClient
type MockPublicIPAddressesClientMockRecorder struct {
	mock *MockPublicIPAddressesClient
}

// NewMockPublicIPAddressesClient creates a new mock instance
func NewMockPublicIPAddressesClient(ctrl *gomock.Controller) *MockPublicIPAddressesClient {
	mock := &MockPublicIPAddressesClient{ctrl: ctrl}
	mock.recorder = &MockPublicIPAddressesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublicIPAddressesClient) EXPECT() *MockPublicIPAddressesClientMockRecorder {
	return m.recorder
}

// DeleteAndWait mocks base method
func (m *MockPublicIPAddressesClient) DeleteAndWait(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAndWait", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAndWait indicates an expected call of DeleteAndWait
func (mr *MockPublicIPAddressesClientMockRecorder) DeleteAndWait(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAndWait", reflect.TypeOf((*MockPublicIPAddressesClient)(nil).DeleteAndWait), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockPublicIPAddressesClient) List(arg0 context.Context, arg1 string) ([]network.PublicIPAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]network.PublicIPAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockPublicIPAddressesClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPublicIPAddressesClient)(nil).List), arg0, arg1)
}
