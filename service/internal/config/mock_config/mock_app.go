// Code generated by mockery v2.12.0. DO NOT EDIT.

package mock_config

import (
	config "simon/mall/service/internal/config"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockAppConfig is an autogenerated mock type for the IAppConfig type
type MockAppConfig struct {
	mock.Mock
}

type MockAppConfig_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAppConfig) EXPECT() *MockAppConfig_Expecter {
	return &MockAppConfig_Expecter{mock: &_m.Mock}
}

// GetAppLogConfig provides a mock function with given fields:
func (_m *MockAppConfig) GetAppLogConfig() config.LogConfig {
	ret := _m.Called()

	var r0 config.LogConfig
	if rf, ok := ret.Get(0).(func() config.LogConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.LogConfig)
	}

	return r0
}

// MockAppConfig_GetAppLogConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAppLogConfig'
type MockAppConfig_GetAppLogConfig_Call struct {
	*mock.Call
}

// GetAppLogConfig is a helper method to define mock.On call
func (_e *MockAppConfig_Expecter) GetAppLogConfig() *MockAppConfig_GetAppLogConfig_Call {
	return &MockAppConfig_GetAppLogConfig_Call{Call: _e.mock.On("GetAppLogConfig")}
}

func (_c *MockAppConfig_GetAppLogConfig_Call) Run(run func()) *MockAppConfig_GetAppLogConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAppConfig_GetAppLogConfig_Call) Return(_a0 config.LogConfig) *MockAppConfig_GetAppLogConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetGinConfig provides a mock function with given fields:
func (_m *MockAppConfig) GetGinConfig() config.GinConfig {
	ret := _m.Called()

	var r0 config.GinConfig
	if rf, ok := ret.Get(0).(func() config.GinConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.GinConfig)
	}

	return r0
}

// MockAppConfig_GetGinConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGinConfig'
type MockAppConfig_GetGinConfig_Call struct {
	*mock.Call
}

// GetGinConfig is a helper method to define mock.On call
func (_e *MockAppConfig_Expecter) GetGinConfig() *MockAppConfig_GetGinConfig_Call {
	return &MockAppConfig_GetGinConfig_Call{Call: _e.mock.On("GetGinConfig")}
}

func (_c *MockAppConfig_GetGinConfig_Call) Run(run func()) *MockAppConfig_GetGinConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAppConfig_GetGinConfig_Call) Return(_a0 config.GinConfig) *MockAppConfig_GetGinConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetLocalCacheConfig provides a mock function with given fields:
func (_m *MockAppConfig) GetLocalCacheConfig() config.LocalCacheConfig {
	ret := _m.Called()

	var r0 config.LocalCacheConfig
	if rf, ok := ret.Get(0).(func() config.LocalCacheConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.LocalCacheConfig)
	}

	return r0
}

// MockAppConfig_GetLocalCacheConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLocalCacheConfig'
type MockAppConfig_GetLocalCacheConfig_Call struct {
	*mock.Call
}

// GetLocalCacheConfig is a helper method to define mock.On call
func (_e *MockAppConfig_Expecter) GetLocalCacheConfig() *MockAppConfig_GetLocalCacheConfig_Call {
	return &MockAppConfig_GetLocalCacheConfig_Call{Call: _e.mock.On("GetLocalCacheConfig")}
}

func (_c *MockAppConfig_GetLocalCacheConfig_Call) Run(run func()) *MockAppConfig_GetLocalCacheConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAppConfig_GetLocalCacheConfig_Call) Return(_a0 config.LocalCacheConfig) *MockAppConfig_GetLocalCacheConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetMySQLConfig provides a mock function with given fields:
func (_m *MockAppConfig) GetMySQLConfig() config.MySQLConfig {
	ret := _m.Called()

	var r0 config.MySQLConfig
	if rf, ok := ret.Get(0).(func() config.MySQLConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.MySQLConfig)
	}

	return r0
}

// MockAppConfig_GetMySQLConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMySQLConfig'
type MockAppConfig_GetMySQLConfig_Call struct {
	*mock.Call
}

// GetMySQLConfig is a helper method to define mock.On call
func (_e *MockAppConfig_Expecter) GetMySQLConfig() *MockAppConfig_GetMySQLConfig_Call {
	return &MockAppConfig_GetMySQLConfig_Call{Call: _e.mock.On("GetMySQLConfig")}
}

func (_c *MockAppConfig_GetMySQLConfig_Call) Run(run func()) *MockAppConfig_GetMySQLConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAppConfig_GetMySQLConfig_Call) Return(_a0 config.MySQLConfig) *MockAppConfig_GetMySQLConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetServerConfig provides a mock function with given fields:
func (_m *MockAppConfig) GetServerConfig() config.ServerConfig {
	ret := _m.Called()

	var r0 config.ServerConfig
	if rf, ok := ret.Get(0).(func() config.ServerConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.ServerConfig)
	}

	return r0
}

// MockAppConfig_GetServerConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetServerConfig'
type MockAppConfig_GetServerConfig_Call struct {
	*mock.Call
}

// GetServerConfig is a helper method to define mock.On call
func (_e *MockAppConfig_Expecter) GetServerConfig() *MockAppConfig_GetServerConfig_Call {
	return &MockAppConfig_GetServerConfig_Call{Call: _e.mock.On("GetServerConfig")}
}

func (_c *MockAppConfig_GetServerConfig_Call) Run(run func()) *MockAppConfig_GetServerConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAppConfig_GetServerConfig_Call) Return(_a0 config.ServerConfig) *MockAppConfig_GetServerConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewMockAppConfig creates a new instance of MockAppConfig. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAppConfig(t testing.TB) *MockAppConfig {
	mock := &MockAppConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
