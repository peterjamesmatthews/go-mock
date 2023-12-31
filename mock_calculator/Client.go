// Code generated by MockGen. DO NOT EDIT.
// Source: ./calculator/Calculatorer.go
//
// Generated by this command:
//
//	mockgen -source=./calculator/Calculatorer.go -destination=./mock_calculator/Client.go
//
// Package mock_calculator is a generated GoMock package.
package mock_calculator

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCalculatorer is a mock of Calculatorer interface.
type MockCalculatorer struct {
	ctrl     *gomock.Controller
	recorder *MockCalculatorerMockRecorder
}

// MockCalculatorerMockRecorder is the mock recorder for MockCalculatorer.
type MockCalculatorerMockRecorder struct {
	mock *MockCalculatorer
}

// NewMockCalculatorer creates a new mock instance.
func NewMockCalculatorer(ctrl *gomock.Controller) *MockCalculatorer {
	mock := &MockCalculatorer{ctrl: ctrl}
	mock.recorder = &MockCalculatorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCalculatorer) EXPECT() *MockCalculatorerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCalculatorer) Add(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockCalculatorerMockRecorder) Add(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCalculatorer)(nil).Add), arg0, arg1)
}

// Divide mocks base method.
func (m *MockCalculatorer) Divide(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Divide", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// Divide indicates an expected call of Divide.
func (mr *MockCalculatorerMockRecorder) Divide(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Divide", reflect.TypeOf((*MockCalculatorer)(nil).Divide), arg0, arg1)
}

// Multiply mocks base method.
func (m *MockCalculatorer) Multiply(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Multiply", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// Multiply indicates an expected call of Multiply.
func (mr *MockCalculatorerMockRecorder) Multiply(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Multiply", reflect.TypeOf((*MockCalculatorer)(nil).Multiply), arg0, arg1)
}

// Subtract mocks base method.
func (m *MockCalculatorer) Subtract(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subtract", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// Subtract indicates an expected call of Subtract.
func (mr *MockCalculatorerMockRecorder) Subtract(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subtract", reflect.TypeOf((*MockCalculatorer)(nil).Subtract), arg0, arg1)
}
