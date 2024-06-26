// Code generated by MockGen. DO NOT EDIT.
// Source: http_request_handler.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHttpRequestHandler is a mock of HttpRequestHandler interface.
type MockHttpRequestHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHttpRequestHandlerMockRecorder
}

// MockHttpRequestHandlerMockRecorder is the mock recorder for MockHttpRequestHandler.
type MockHttpRequestHandlerMockRecorder struct {
	mock *MockHttpRequestHandler
}

// NewMockHttpRequestHandler creates a new mock instance.
func NewMockHttpRequestHandler(ctrl *gomock.Controller) *MockHttpRequestHandler {
	mock := &MockHttpRequestHandler{ctrl: ctrl}
	mock.recorder = &MockHttpRequestHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpRequestHandler) EXPECT() *MockHttpRequestHandlerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockHttpRequestHandler) Get(url string, repsonseModel interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", url, repsonseModel)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockHttpRequestHandlerMockRecorder) Get(url, repsonseModel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHttpRequestHandler)(nil).Get), url, repsonseModel)
}
