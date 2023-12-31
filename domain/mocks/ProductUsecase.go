// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/jordanlanch/docucenter-test/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ProductUsecase is an autogenerated mock type for the ProductUsecase type
type ProductUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: p
func (_m *ProductUsecase) Create(p *domain.Product) (*domain.Product, error) {
	ret := _m.Called(p)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(*domain.Product) *domain.Product); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Product) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *ProductUsecase) GetByID(id uuid.UUID) (*domain.Product, error) {
	ret := _m.Called(id)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(uuid.UUID) *domain.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
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

// GetMany provides a mock function with given fields: pagination, filters
func (_m *ProductUsecase) GetMany(pagination *domain.Pagination, filters map[string]interface{}) ([]*domain.Product, error) {
	ret := _m.Called(pagination, filters)

	var r0 []*domain.Product
	if rf, ok := ret.Get(0).(func(*domain.Pagination, map[string]interface{}) []*domain.Product); ok {
		r0 = rf(pagination, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Pagination, map[string]interface{}) error); ok {
		r1 = rf(pagination, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modify provides a mock function with given fields: p, id
func (_m *ProductUsecase) Modify(p *domain.Product, id uuid.UUID) (*domain.Product, error) {
	ret := _m.Called(p, id)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(*domain.Product, uuid.UUID) *domain.Product); ok {
		r0 = rf(p, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Product, uuid.UUID) error); ok {
		r1 = rf(p, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: id
func (_m *ProductUsecase) Remove(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewProductUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductUsecase creates a new instance of ProductUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductUsecase(t mockConstructorTestingTNewProductUsecase) *ProductUsecase {
	mock := &ProductUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
