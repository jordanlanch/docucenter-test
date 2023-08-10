// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/jordanlanch/docucenter-test/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// WarehousePortUsecase is an autogenerated mock type for the WarehousePortUsecase type
type WarehousePortUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: wp
func (_m *WarehousePortUsecase) Create(wp *domain.WarehousesAndPorts) (*domain.WarehousesAndPorts, error) {
	ret := _m.Called(wp)

	var r0 *domain.WarehousesAndPorts
	if rf, ok := ret.Get(0).(func(*domain.WarehousesAndPorts) *domain.WarehousesAndPorts); ok {
		r0 = rf(wp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.WarehousesAndPorts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.WarehousesAndPorts) error); ok {
		r1 = rf(wp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *WarehousePortUsecase) GetByID(id uuid.UUID) (*domain.WarehousesAndPorts, error) {
	ret := _m.Called(id)

	var r0 *domain.WarehousesAndPorts
	if rf, ok := ret.Get(0).(func(uuid.UUID) *domain.WarehousesAndPorts); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.WarehousesAndPorts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMany provides a mock function with given fields: pagination
func (_m *WarehousePortUsecase) GetMany(pagination *domain.Pagination) ([]*domain.WarehousesAndPorts, error) {
	ret := _m.Called(pagination)

	var r0 []*domain.WarehousesAndPorts
	if rf, ok := ret.Get(0).(func(*domain.Pagination) []*domain.WarehousesAndPorts); ok {
		r0 = rf(pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.WarehousesAndPorts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Pagination) error); ok {
		r1 = rf(pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modify provides a mock function with given fields: wp, id
func (_m *WarehousePortUsecase) Modify(wp *domain.WarehousesAndPorts, id uuid.UUID) (*domain.WarehousesAndPorts, error) {
	ret := _m.Called(wp, id)

	var r0 *domain.WarehousesAndPorts
	if rf, ok := ret.Get(0).(func(*domain.WarehousesAndPorts, uuid.UUID) *domain.WarehousesAndPorts); ok {
		r0 = rf(wp, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.WarehousesAndPorts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.WarehousesAndPorts, uuid.UUID) error); ok {
		r1 = rf(wp, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: id
func (_m *WarehousePortUsecase) Remove(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewWarehousePortUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewWarehousePortUsecase creates a new instance of WarehousePortUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWarehousePortUsecase(t mockConstructorTestingTNewWarehousePortUsecase) *WarehousePortUsecase {
	mock := &WarehousePortUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
