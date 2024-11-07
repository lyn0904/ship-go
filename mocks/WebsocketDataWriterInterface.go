// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	api "github.com/lyn0904/ship-go/api"
	mock "github.com/stretchr/testify/mock"
)

// WebsocketDataWriterInterface is an autogenerated mock type for the WebsocketDataWriterInterface type
type WebsocketDataWriterInterface struct {
	mock.Mock
}

type WebsocketDataWriterInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *WebsocketDataWriterInterface) EXPECT() *WebsocketDataWriterInterface_Expecter {
	return &WebsocketDataWriterInterface_Expecter{mock: &_m.Mock}
}

// CloseDataConnection provides a mock function with given fields: closeCode, reason
func (_m *WebsocketDataWriterInterface) CloseDataConnection(closeCode int, reason string) {
	_m.Called(closeCode, reason)
}

// WebsocketDataWriterInterface_CloseDataConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseDataConnection'
type WebsocketDataWriterInterface_CloseDataConnection_Call struct {
	*mock.Call
}

// CloseDataConnection is a helper method to define mock.On call
//   - closeCode int
//   - reason string
func (_e *WebsocketDataWriterInterface_Expecter) CloseDataConnection(closeCode interface{}, reason interface{}) *WebsocketDataWriterInterface_CloseDataConnection_Call {
	return &WebsocketDataWriterInterface_CloseDataConnection_Call{Call: _e.mock.On("CloseDataConnection", closeCode, reason)}
}

func (_c *WebsocketDataWriterInterface_CloseDataConnection_Call) Run(run func(closeCode int, reason string)) *WebsocketDataWriterInterface_CloseDataConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(string))
	})
	return _c
}

func (_c *WebsocketDataWriterInterface_CloseDataConnection_Call) Return() *WebsocketDataWriterInterface_CloseDataConnection_Call {
	_c.Call.Return()
	return _c
}

func (_c *WebsocketDataWriterInterface_CloseDataConnection_Call) RunAndReturn(run func(int, string)) *WebsocketDataWriterInterface_CloseDataConnection_Call {
	_c.Call.Return(run)
	return _c
}

// InitDataProcessing provides a mock function with given fields: _a0
func (_m *WebsocketDataWriterInterface) InitDataProcessing(_a0 api.WebsocketDataReaderInterface) {
	_m.Called(_a0)
}

// WebsocketDataWriterInterface_InitDataProcessing_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitDataProcessing'
type WebsocketDataWriterInterface_InitDataProcessing_Call struct {
	*mock.Call
}

// InitDataProcessing is a helper method to define mock.On call
//   - _a0 api.WebsocketDataReaderInterface
func (_e *WebsocketDataWriterInterface_Expecter) InitDataProcessing(_a0 interface{}) *WebsocketDataWriterInterface_InitDataProcessing_Call {
	return &WebsocketDataWriterInterface_InitDataProcessing_Call{Call: _e.mock.On("InitDataProcessing", _a0)}
}

func (_c *WebsocketDataWriterInterface_InitDataProcessing_Call) Run(run func(_a0 api.WebsocketDataReaderInterface)) *WebsocketDataWriterInterface_InitDataProcessing_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.WebsocketDataReaderInterface))
	})
	return _c
}

func (_c *WebsocketDataWriterInterface_InitDataProcessing_Call) Return() *WebsocketDataWriterInterface_InitDataProcessing_Call {
	_c.Call.Return()
	return _c
}

func (_c *WebsocketDataWriterInterface_InitDataProcessing_Call) RunAndReturn(run func(api.WebsocketDataReaderInterface)) *WebsocketDataWriterInterface_InitDataProcessing_Call {
	_c.Call.Return(run)
	return _c
}

// IsDataConnectionClosed provides a mock function with given fields:
func (_m *WebsocketDataWriterInterface) IsDataConnectionClosed() (bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsDataConnectionClosed")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WebsocketDataWriterInterface_IsDataConnectionClosed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsDataConnectionClosed'
type WebsocketDataWriterInterface_IsDataConnectionClosed_Call struct {
	*mock.Call
}

// IsDataConnectionClosed is a helper method to define mock.On call
func (_e *WebsocketDataWriterInterface_Expecter) IsDataConnectionClosed() *WebsocketDataWriterInterface_IsDataConnectionClosed_Call {
	return &WebsocketDataWriterInterface_IsDataConnectionClosed_Call{Call: _e.mock.On("IsDataConnectionClosed")}
}

func (_c *WebsocketDataWriterInterface_IsDataConnectionClosed_Call) Run(run func()) *WebsocketDataWriterInterface_IsDataConnectionClosed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WebsocketDataWriterInterface_IsDataConnectionClosed_Call) Return(_a0 bool, _a1 error) *WebsocketDataWriterInterface_IsDataConnectionClosed_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WebsocketDataWriterInterface_IsDataConnectionClosed_Call) RunAndReturn(run func() (bool, error)) *WebsocketDataWriterInterface_IsDataConnectionClosed_Call {
	_c.Call.Return(run)
	return _c
}

// WriteMessageToWebsocketConnection provides a mock function with given fields: _a0
func (_m *WebsocketDataWriterInterface) WriteMessageToWebsocketConnection(_a0 []byte) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for WriteMessageToWebsocketConnection")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WebsocketDataWriterInterface_WriteMessageToWebsocketConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteMessageToWebsocketConnection'
type WebsocketDataWriterInterface_WriteMessageToWebsocketConnection_Call struct {
	*mock.Call
}

// WriteMessageToWebsocketConnection is a helper method to define mock.On call
//   - _a0 []byte
func (_e *WebsocketDataWriterInterface_Expecter) WriteMessageToWebsocketConnection(_a0 interface{}) *WebsocketDataWriterInterface_WriteMessageToWebsocketConnection_Call {
	return &WebsocketDataWriterInterface_WriteMessageToWebsocketConnection_Call{Call: _e.mock.On("WriteMessageToWebsocketConnection", _a0)}
}

func (_c *WebsocketDataWriterInterface_WriteMessageToWebsocketConnection_Call) Run(run func(_a0 []byte)) *WebsocketDataWriterInterface_WriteMessageToWebsocketConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *WebsocketDataWriterInterface_WriteMessageToWebsocketConnection_Call) Return(_a0 error) *WebsocketDataWriterInterface_WriteMessageToWebsocketConnection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WebsocketDataWriterInterface_WriteMessageToWebsocketConnection_Call) RunAndReturn(run func([]byte) error) *WebsocketDataWriterInterface_WriteMessageToWebsocketConnection_Call {
	_c.Call.Return(run)
	return _c
}

// NewWebsocketDataWriterInterface creates a new instance of WebsocketDataWriterInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWebsocketDataWriterInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *WebsocketDataWriterInterface {
	mock := &WebsocketDataWriterInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
