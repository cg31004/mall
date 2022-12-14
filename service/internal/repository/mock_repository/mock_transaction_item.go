// Code generated by mockery v2.12.0. DO NOT EDIT.

package mock_repository

import (
	context "context"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	po "simon/mall/service/internal/model/po"

	testing "testing"
)

// MockTxnItemRepo is an autogenerated mock type for the ITxnItemRepo type
type MockTxnItemRepo struct {
	mock.Mock
}

type MockTxnItemRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTxnItemRepo) EXPECT() *MockTxnItemRepo_Expecter {
	return &MockTxnItemRepo_Expecter{mock: &_m.Mock}
}

// GetList provides a mock function with given fields: ctx, db, cond
func (_m *MockTxnItemRepo) GetList(ctx context.Context, db *gorm.DB, cond *po.GetTxnItemListCond) ([]*po.TransactionItem, error) {
	ret := _m.Called(ctx, db, cond)

	var r0 []*po.TransactionItem
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *po.GetTxnItemListCond) []*po.TransactionItem); ok {
		r0 = rf(ctx, db, cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*po.TransactionItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, *po.GetTxnItemListCond) error); ok {
		r1 = rf(ctx, db, cond)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTxnItemRepo_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockTxnItemRepo_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//  - ctx context.Context
//  - db *gorm.DB
//  - cond *po.GetTxnItemListCond
func (_e *MockTxnItemRepo_Expecter) GetList(ctx interface{}, db interface{}, cond interface{}) *MockTxnItemRepo_GetList_Call {
	return &MockTxnItemRepo_GetList_Call{Call: _e.mock.On("GetList", ctx, db, cond)}
}

func (_c *MockTxnItemRepo_GetList_Call) Run(run func(ctx context.Context, db *gorm.DB, cond *po.GetTxnItemListCond)) *MockTxnItemRepo_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*gorm.DB), args[2].(*po.GetTxnItemListCond))
	})
	return _c
}

func (_c *MockTxnItemRepo_GetList_Call) Return(_a0 []*po.TransactionItem, _a1 error) *MockTxnItemRepo_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Insert provides a mock function with given fields: ctx, db, txnItem
func (_m *MockTxnItemRepo) Insert(ctx context.Context, db *gorm.DB, txnItem []*po.TransactionItem) error {
	ret := _m.Called(ctx, db, txnItem)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, []*po.TransactionItem) error); ok {
		r0 = rf(ctx, db, txnItem)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTxnItemRepo_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type MockTxnItemRepo_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//  - ctx context.Context
//  - db *gorm.DB
//  - txnItem []*po.TransactionItem
func (_e *MockTxnItemRepo_Expecter) Insert(ctx interface{}, db interface{}, txnItem interface{}) *MockTxnItemRepo_Insert_Call {
	return &MockTxnItemRepo_Insert_Call{Call: _e.mock.On("Insert", ctx, db, txnItem)}
}

func (_c *MockTxnItemRepo_Insert_Call) Run(run func(ctx context.Context, db *gorm.DB, txnItem []*po.TransactionItem)) *MockTxnItemRepo_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*gorm.DB), args[2].([]*po.TransactionItem))
	})
	return _c
}

func (_c *MockTxnItemRepo_Insert_Call) Return(_a0 error) *MockTxnItemRepo_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewMockTxnItemRepo creates a new instance of MockTxnItemRepo. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTxnItemRepo(t testing.TB) *MockTxnItemRepo {
	mock := &MockTxnItemRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
