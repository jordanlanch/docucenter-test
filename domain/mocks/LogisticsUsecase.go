// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/jordanlanch/docucenter-test/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// LogisticsUsecase is an autogenerated mock type for the LogisticsUsecase type
type LogisticsUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ll
func (_m *LogisticsUsecase) Create(ll *domain.Logistics) (*domain.Logistics, error) {
	ret := _m.Called(ll)

	var r0 *domain.Logistics
	if rf, ok := ret.Get(0).(func(*domain.Logistics) *domain.Logistics); ok {
		r0 = rf(ll)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Logistics)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Logistics) error); ok {
		r1 = rf(ll)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *LogisticsUsecase) GetByID(id uuid.UUID) (*domain.Logistics, error) {
	ret := _m.Called(id)

	var r0 *domain.Logistics
	if rf, ok := ret.Get(0).(func(uuid.UUID) *domain.Logistics); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Logistics)
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
func (_m *LogisticsUsecase) GetMany(pagination *domain.Pagination) ([]*domain.Logistics, error) {
	ret := _m.Called(pagination)

	var r0 []*domain.Logistics
	if rf, ok := ret.Get(0).(func(*domain.Pagination) []*domain.Logistics); ok {
		r0 = rf(pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Logistics)
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

// Modify provides a mock function with given fields: ll, id
func (_m *LogisticsUsecase) Modify(ll *domain.Logistics, id uuid.UUID) (*domain.Logistics, error) {
	ret := _m.Called(ll, id)

	var r0 *domain.Logistics
	if rf, ok := ret.Get(0).(func(*domain.Logistics, uuid.UUID) *domain.Logistics); ok {
		r0 = rf(ll, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Logistics)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Logistics, uuid.UUID) error); ok {
		r1 = rf(ll, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: id
func (_m *LogisticsUsecase) Remove(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewLogisticsUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogisticsUsecase creates a new instance of LogisticsUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogisticsUsecase(t mockConstructorTestingTNewLogisticsUsecase) *LogisticsUsecase {
	mock := &LogisticsUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
