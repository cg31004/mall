// Code generated by mockery v2.12.0. DO NOT EDIT.

package mock_repository

import (
	context "context"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	po "simon/mall/service/internal/model/po"

	testing "testing"
)

// MockPaymentRepo is an autogenerated mock type for the IProductRepo type
type MockPaymentRepo struct {
	mock.Mock
}

type MockPaymentRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPaymentRepo) EXPECT() *MockPaymentRepo_Expecter {
	return &MockPaymentRepo_Expecter{mock: &_m.Mock}
}

// First provides a mock function with given fields: ctx, db, id
func (_m *MockPaymentRepo) First(ctx context.Context, db *gorm.DB, id string) (*po.Product, error) {
	ret := _m.Called(ctx, db, id)

	var r0 *po.Product
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, string) *po.Product); ok {
		r0 = rf(ctx, db, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*po.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, string) error); ok {
		r1 = rf(ctx, db, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPaymentRepo_First_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'First'
type MockPaymentRepo_First_Call struct {
	*mock.Call
}

// First is a helper method to define mock.On call
//  - ctx context.Context
//  - db *gorm.DB
//  - id string
func (_e *MockPaymentRepo_Expecter) First(ctx interface{}, db interface{}, id interface{}) *MockPaymentRepo_First_Call {
	return &MockPaymentRepo_First_Call{Call: _e.mock.On("First", ctx, db, id)}
}

func (_c *MockPaymentRepo_First_Call) Run(run func(ctx context.Context, db *gorm.DB, id string)) *MockPaymentRepo_First_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*gorm.DB), args[2].(string))
	})
	return _c
}

func (_c *MockPaymentRepo_First_Call) Return(_a0 *po.Product, _a1 error) *MockPaymentRepo_First_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetList provides a mock function with given fields: ctx, db, cond, pager
func (_m *MockPaymentRepo) GetList(ctx context.Context, db *gorm.DB, cond *po.ProductSearch, pager *po.Pager) ([]*po.Product, error) {
	ret := _m.Called(ctx, db, cond, pager)

	var r0 []*po.Product
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *po.ProductSearch, *po.Pager) []*po.Product); ok {
		r0 = rf(ctx, db, cond, pager)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*po.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, *po.ProductSearch, *po.Pager) error); ok {
		r1 = rf(ctx, db, cond, pager)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPaymentRepo_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockPaymentRepo_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//  - ctx context.Context
//  - db *gorm.DB
//  - cond *po.ProductSearch
//  - pager *po.Pager
func (_e *MockPaymentRepo_Expecter) GetList(ctx interface{}, db interface{}, cond interface{}, pager interface{}) *MockPaymentRepo_GetList_Call {
	return &MockPaymentRepo_GetList_Call{Call: _e.mock.On("GetList", ctx, db, cond, pager)}
}

func (_c *MockPaymentRepo_GetList_Call) Run(run func(ctx context.Context, db *gorm.DB, cond *po.ProductSearch, pager *po.Pager)) *MockPaymentRepo_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*gorm.DB), args[2].(*po.ProductSearch), args[3].(*po.Pager))
	})
	return _c
}

func (_c *MockPaymentRepo_GetList_Call) Return(_a0 []*po.Product, _a1 error) *MockPaymentRepo_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetListPager provides a mock function with given fields: ctx, db, cond, pager
func (_m *MockPaymentRepo) GetListPager(ctx context.Context, db *gorm.DB, cond *po.ProductSearch, pager *po.Pager) (*po.PagingResult, error) {
	ret := _m.Called(ctx, db, cond, pager)

	var r0 *po.PagingResult
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *po.ProductSearch, *po.Pager) *po.PagingResult); ok {
		r0 = rf(ctx, db, cond, pager)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*po.PagingResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, *po.ProductSearch, *po.Pager) error); ok {
		r1 = rf(ctx, db, cond, pager)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPaymentRepo_GetListPager_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListPager'
type MockPaymentRepo_GetListPager_Call struct {
	*mock.Call
}

// GetListPager is a helper method to define mock.On call
//  - ctx context.Context
//  - db *gorm.DB
//  - cond *po.ProductSearch
//  - pager *po.Pager
func (_e *MockPaymentRepo_Expecter) GetListPager(ctx interface{}, db interface{}, cond interface{}, pager interface{}) *MockPaymentRepo_GetListPager_Call {
	return &MockPaymentRepo_GetListPager_Call{Call: _e.mock.On("GetListPager", ctx, db, cond, pager)}
}

func (_c *MockPaymentRepo_GetListPager_Call) Run(run func(ctx context.Context, db *gorm.DB, cond *po.ProductSearch, pager *po.Pager)) *MockPaymentRepo_GetListPager_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*gorm.DB), args[2].(*po.ProductSearch), args[3].(*po.Pager))
	})
	return _c
}

func (_c *MockPaymentRepo_GetListPager_Call) Return(_a0 *po.PagingResult, _a1 error) *MockPaymentRepo_GetListPager_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// NewMockPaymentRepo creates a new instance of MockPaymentRepo. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPaymentRepo(t testing.TB) *MockPaymentRepo {
	mock := &MockPaymentRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}