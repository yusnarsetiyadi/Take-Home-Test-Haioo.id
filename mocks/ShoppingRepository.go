// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	shopping "haiooId/shoppingChart/features/shopping"

	mock "github.com/stretchr/testify/mock"
)

// ShoppingRepository is an autogenerated mock type for the RepositoryInterface type
type ShoppingRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *ShoppingRepository) Create(input shopping.ShoppingCore) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(shopping.ShoppingCore) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: product_code
func (_m *ShoppingRepository) Delete(product_code uint) error {
	ret := _m.Called(product_code)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(product_code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindProduct provides a mock function with given fields: product
func (_m *ShoppingRepository) FindProduct(product string) (shopping.ShoppingCore, error) {
	ret := _m.Called(product)

	var r0 shopping.ShoppingCore
	if rf, ok := ret.Get(0).(func(string) shopping.ShoppingCore); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Get(0).(shopping.ShoppingCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: product_name, qty
func (_m *ShoppingRepository) Get(product_name string, qty string) ([]shopping.ShoppingCore, error) {
	ret := _m.Called(product_name, qty)

	var r0 []shopping.ShoppingCore
	if rf, ok := ret.Get(0).(func(string, string) []shopping.ShoppingCore); ok {
		r0 = rf(product_name, qty)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]shopping.ShoppingCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(product_name, qty)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: input, product_code
func (_m *ShoppingRepository) Update(input shopping.ShoppingCore, product_code uint) error {
	ret := _m.Called(input, product_code)

	var r0 error
	if rf, ok := ret.Get(0).(func(shopping.ShoppingCore, uint) error); ok {
		r0 = rf(input, product_code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewShoppingRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewShoppingRepository creates a new instance of ShoppingRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewShoppingRepository(t mockConstructorTestingTNewShoppingRepository) *ShoppingRepository {
	mock := &ShoppingRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
