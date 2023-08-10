// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/jordanlanch/docucenter-test/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// DiscountUsecase is an autogenerated mock type for the DiscountUsecase type
type DiscountUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: d
func (_m *DiscountUsecase) Create(d *domain.Discount) (*domain.Discount, error) {
	ret := _m.Called(d)

	var r0 *domain.Discount
	if rf, ok := ret.Get(0).(func(*domain.Discount) *domain.Discount); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Discount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Discount) error); ok {
		r1 = rf(d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTypeAndQuantity provides a mock function with given fields: dtype, quantity
func (_m *DiscountUsecase) GetByTypeAndQuantity(dtype domain.LogisticsType, quantity int) (*domain.Discount, error) {
	ret := _m.Called(dtype, quantity)

	var r0 *domain.Discount
	if rf, ok := ret.Get(0).(func(domain.LogisticsType, int) *domain.Discount); ok {
		r0 = rf(dtype, quantity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Discount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.LogisticsType, int) error); ok {
		r1 = rf(dtype, quantity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMany provides a mock function with given fields: pagination
func (_m *DiscountUsecase) GetMany(pagination *domain.Pagination) ([]*domain.Discount, error) {
	ret := _m.Called(pagination)

	var r0 []*domain.Discount
	if rf, ok := ret.Get(0).(func(*domain.Pagination) []*domain.Discount); ok {
		r0 = rf(pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Discount)
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

// Modify provides a mock function with given fields: d
func (_m *DiscountUsecase) Modify(d *domain.Discount) (*domain.Discount, error) {
	ret := _m.Called(d)

	var r0 *domain.Discount
	if rf, ok := ret.Get(0).(func(*domain.Discount) *domain.Discount); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Discount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Discount) error); ok {
		r1 = rf(d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: id
func (_m *DiscountUsecase) Remove(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDiscountUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewDiscountUsecase creates a new instance of DiscountUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDiscountUsecase(t mockConstructorTestingTNewDiscountUsecase) *DiscountUsecase {
	mock := &DiscountUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
