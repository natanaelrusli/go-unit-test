// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	entity "github.com/natanaelrusli/go-unit-test/entity"
	mock "github.com/stretchr/testify/mock"
)

// CategoryRepository is an autogenerated mock type for the CategoryRepository type
type CategoryRepository struct {
	mock.Mock
}

// FindById provides a mock function with given fields: id
func (_m *CategoryRepository) FindById(id string) *entity.Category {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *entity.Category
	if rf, ok := ret.Get(0).(func(string) *entity.Category); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Category)
		}
	}

	return r0
}

// NewCategoryRepository creates a new instance of CategoryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCategoryRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CategoryRepository {
	mock := &CategoryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
