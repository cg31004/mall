// Code generated by mockery v2.12.0. DO NOT EDIT.

package mock_repository

import (
	context "context"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	po "simon/mall/service/internal/model/po"

	testing "testing"
)

// MockTxnRepo is an autogenerated mock type for the ITxnRepo type
type MockTxnRepo struct {
	mock.Mock
}

type MockTxnRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTxnRepo) EXPECT() *MockTxnRepo_Expecter {
	return &MockTxnRepo_Expecter{mock: &_m.Mock}
}

// GetList provides a mock function with given fields: ctx, db, cond
func (_m *MockTxnRepo) GetList(ctx context.Context, db *gorm.DB, cond *po.TransactionSearch) ([]*po.Transaction, error) {
	ret := _m.Called(ctx, db, cond)

	var r0 []*po.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *po.TransactionSearch) []*po.Transaction); ok {
		r0 = rf(ctx, db, cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*po.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, *po.TransactionSearch) error); ok {
		r1 = rf(ctx, db, cond)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTxnRepo_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockTxnRepo_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//  - ctx context.Context
//  - db *gorm.DB
//  - cond *po.TransactionSearch
func (_e *MockTxnRepo_Expecter) GetList(ctx interface{}, db interface{}, cond interface{}) *MockTxnRepo_GetList_Call {
	return &MockTxnRepo_GetList_Call{Call: _e.mock.On("GetList", ctx, db, cond)}
}

func (_c *MockTxnRepo_GetList_Call) Run(run func(ctx context.Context, db *gorm.DB, cond *po.TransactionSearch)) *MockTxnRepo_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*gorm.DB), args[2].(*po.TransactionSearch))
	})
	return _c
}

func (_c *MockTxnRepo_GetList_Call) Return(_a0 []*po.Transaction, _a1 error) *MockTxnRepo_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Insert provides a mock function with given fields: ctx, db, txnItem
func (_m *MockTxnRepo) Insert(ctx context.Context, db *gorm.DB, txnItem *po.Transaction) error {
	ret := _m.Called(ctx, db, txnItem)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *po.Transaction) error); ok {
		r0 = rf(ctx, db, txnItem)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTxnRepo_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type MockTxnRepo_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//  - ctx context.Context
//  - db *gorm.DB
//  - txnItem *po.Transaction
func (_e *MockTxnRepo_Expecter) Insert(ctx interface{}, db interface{}, txnItem interface{}) *MockTxnRepo_Insert_Call {
	return &MockTxnRepo_Insert_Call{Call: _e.mock.On("Insert", ctx, db, txnItem)}
}

func (_c *MockTxnRepo_Insert_Call) Run(run func(ctx context.Context, db *gorm.DB, txnItem *po.Transaction)) *MockTxnRepo_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*gorm.DB), args[2].(*po.Transaction))
	})
	return _c
}

func (_c *MockTxnRepo_Insert_Call) Return(_a0 error) *MockTxnRepo_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewMockTxnRepo creates a new instance of MockTxnRepo. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTxnRepo(t testing.TB) *MockTxnRepo {
	mock := &MockTxnRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
