// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/license/client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// IsLicenseValid mocks base method
func (m *MockClient) IsLicenseValid() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLicenseValid")
	ret0, _ := ret[0].(error)
	return ret0
}

// IsLicenseValid indicates an expected call of IsLicenseValid
func (mr *MockClientMockRecorder) IsLicenseValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLicenseValid", reflect.TypeOf((*MockClient)(nil).IsLicenseValid))
}
