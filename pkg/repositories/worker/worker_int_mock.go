// Code generated by MockGen. DO NOT EDIT.
// Source: worker_int.go

// Package worker is a generated GoMock package.
package worker

import (
	reflect "reflect"
	handler "relap/pkg/repositories/handler"

	gomock "github.com/golang/mock/gomock"
)

// MockInt is a mock of Int interface
type MockInt struct {
	ctrl     *gomock.Controller
	recorder *MockIntMockRecorder
}

// MockIntMockRecorder is the mock recorder for MockInt
type MockIntMockRecorder struct {
	mock *MockInt
}

// NewMockInt creates a new mock instance
func NewMockInt(ctrl *gomock.Controller) *MockInt {
	mock := &MockInt{ctrl: ctrl}
	mock.recorder = &MockIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInt) EXPECT() *MockIntMockRecorder {
	return m.recorder
}

// FetchPage mocks base method
func (m *MockInt) FetchPage(url string, categories []string) (*handler.ResultData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchPage", url, categories)
	ret0, _ := ret[0].(*handler.ResultData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchPage indicates an expected call of FetchPage
func (mr *MockIntMockRecorder) FetchPage(url, categories interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPage", reflect.TypeOf((*MockInt)(nil).FetchPage), url, categories)
}
