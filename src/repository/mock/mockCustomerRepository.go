// Code generated by MockGen. DO NOT EDIT.
// Source: customerRepository.go
//
// Generated by this command:
//
//	mockgen -source=customerRepository.go -destination=../../repository/mock/mockCustomerRepository.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	entity "github.com/overusevery/golang-echo-practice2/src/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockCustomerRepository is a mock of CustomerRepository interface.
type MockCustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerRepositoryMockRecorder
}

// MockCustomerRepositoryMockRecorder is the mock recorder for MockCustomerRepository.
type MockCustomerRepositoryMockRecorder struct {
	mock *MockCustomerRepository
}

// NewMockCustomerRepository creates a new mock instance.
func NewMockCustomerRepository(ctrl *gomock.Controller) *MockCustomerRepository {
	mock := &MockCustomerRepository{ctrl: ctrl}
	mock.recorder = &MockCustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerRepository) EXPECT() *MockCustomerRepositoryMockRecorder {
	return m.recorder
}

// GetCustomer mocks base method.
func (m *MockCustomerRepository) GetCustomer(id int) entity.Customer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomer", id)
	ret0, _ := ret[0].(entity.Customer)
	return ret0
}

// GetCustomer indicates an expected call of GetCustomer.
func (mr *MockCustomerRepositoryMockRecorder) GetCustomer(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomer", reflect.TypeOf((*MockCustomerRepository)(nil).GetCustomer), id)
}
