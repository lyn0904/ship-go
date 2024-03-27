// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/ship-go/api"
	mock "github.com/stretchr/testify/mock"

	model "github.com/enbility/ship-go/model"
)

// ShipConnectionInfoProviderInterface is an autogenerated mock type for the ShipConnectionInfoProviderInterface type
type ShipConnectionInfoProviderInterface struct {
	mock.Mock
}

type ShipConnectionInfoProviderInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ShipConnectionInfoProviderInterface) EXPECT() *ShipConnectionInfoProviderInterface_Expecter {
	return &ShipConnectionInfoProviderInterface_Expecter{mock: &_m.Mock}
}

// AllowWaitingForTrust provides a mock function with given fields: _a0
func (_m *ShipConnectionInfoProviderInterface) AllowWaitingForTrust(_a0 string) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for AllowWaitingForTrust")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ShipConnectionInfoProviderInterface_AllowWaitingForTrust_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllowWaitingForTrust'
type ShipConnectionInfoProviderInterface_AllowWaitingForTrust_Call struct {
	*mock.Call
}

// AllowWaitingForTrust is a helper method to define mock.On call
//   - _a0 string
func (_e *ShipConnectionInfoProviderInterface_Expecter) AllowWaitingForTrust(_a0 interface{}) *ShipConnectionInfoProviderInterface_AllowWaitingForTrust_Call {
	return &ShipConnectionInfoProviderInterface_AllowWaitingForTrust_Call{Call: _e.mock.On("AllowWaitingForTrust", _a0)}
}

func (_c *ShipConnectionInfoProviderInterface_AllowWaitingForTrust_Call) Run(run func(_a0 string)) *ShipConnectionInfoProviderInterface_AllowWaitingForTrust_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_AllowWaitingForTrust_Call) Return(_a0 bool) *ShipConnectionInfoProviderInterface_AllowWaitingForTrust_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_AllowWaitingForTrust_Call) RunAndReturn(run func(string) bool) *ShipConnectionInfoProviderInterface_AllowWaitingForTrust_Call {
	_c.Call.Return(run)
	return _c
}

// HandleConnectionClosed provides a mock function with given fields: _a0, _a1
func (_m *ShipConnectionInfoProviderInterface) HandleConnectionClosed(_a0 api.ShipConnectionInterface, _a1 bool) {
	_m.Called(_a0, _a1)
}

// ShipConnectionInfoProviderInterface_HandleConnectionClosed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleConnectionClosed'
type ShipConnectionInfoProviderInterface_HandleConnectionClosed_Call struct {
	*mock.Call
}

// HandleConnectionClosed is a helper method to define mock.On call
//   - _a0 api.ShipConnectionInterface
//   - _a1 bool
func (_e *ShipConnectionInfoProviderInterface_Expecter) HandleConnectionClosed(_a0 interface{}, _a1 interface{}) *ShipConnectionInfoProviderInterface_HandleConnectionClosed_Call {
	return &ShipConnectionInfoProviderInterface_HandleConnectionClosed_Call{Call: _e.mock.On("HandleConnectionClosed", _a0, _a1)}
}

func (_c *ShipConnectionInfoProviderInterface_HandleConnectionClosed_Call) Run(run func(_a0 api.ShipConnectionInterface, _a1 bool)) *ShipConnectionInfoProviderInterface_HandleConnectionClosed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.ShipConnectionInterface), args[1].(bool))
	})
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_HandleConnectionClosed_Call) Return() *ShipConnectionInfoProviderInterface_HandleConnectionClosed_Call {
	_c.Call.Return()
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_HandleConnectionClosed_Call) RunAndReturn(run func(api.ShipConnectionInterface, bool)) *ShipConnectionInfoProviderInterface_HandleConnectionClosed_Call {
	_c.Call.Return(run)
	return _c
}

// HandleShipHandshakeStateUpdate provides a mock function with given fields: _a0, _a1
func (_m *ShipConnectionInfoProviderInterface) HandleShipHandshakeStateUpdate(_a0 string, _a1 model.ShipState) {
	_m.Called(_a0, _a1)
}

// ShipConnectionInfoProviderInterface_HandleShipHandshakeStateUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleShipHandshakeStateUpdate'
type ShipConnectionInfoProviderInterface_HandleShipHandshakeStateUpdate_Call struct {
	*mock.Call
}

// HandleShipHandshakeStateUpdate is a helper method to define mock.On call
//   - _a0 string
//   - _a1 model.ShipState
func (_e *ShipConnectionInfoProviderInterface_Expecter) HandleShipHandshakeStateUpdate(_a0 interface{}, _a1 interface{}) *ShipConnectionInfoProviderInterface_HandleShipHandshakeStateUpdate_Call {
	return &ShipConnectionInfoProviderInterface_HandleShipHandshakeStateUpdate_Call{Call: _e.mock.On("HandleShipHandshakeStateUpdate", _a0, _a1)}
}

func (_c *ShipConnectionInfoProviderInterface_HandleShipHandshakeStateUpdate_Call) Run(run func(_a0 string, _a1 model.ShipState)) *ShipConnectionInfoProviderInterface_HandleShipHandshakeStateUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(model.ShipState))
	})
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_HandleShipHandshakeStateUpdate_Call) Return() *ShipConnectionInfoProviderInterface_HandleShipHandshakeStateUpdate_Call {
	_c.Call.Return()
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_HandleShipHandshakeStateUpdate_Call) RunAndReturn(run func(string, model.ShipState)) *ShipConnectionInfoProviderInterface_HandleShipHandshakeStateUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// IsAutoAcceptEnabled provides a mock function with given fields:
func (_m *ShipConnectionInfoProviderInterface) IsAutoAcceptEnabled() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsAutoAcceptEnabled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ShipConnectionInfoProviderInterface_IsAutoAcceptEnabled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAutoAcceptEnabled'
type ShipConnectionInfoProviderInterface_IsAutoAcceptEnabled_Call struct {
	*mock.Call
}

// IsAutoAcceptEnabled is a helper method to define mock.On call
func (_e *ShipConnectionInfoProviderInterface_Expecter) IsAutoAcceptEnabled() *ShipConnectionInfoProviderInterface_IsAutoAcceptEnabled_Call {
	return &ShipConnectionInfoProviderInterface_IsAutoAcceptEnabled_Call{Call: _e.mock.On("IsAutoAcceptEnabled")}
}

func (_c *ShipConnectionInfoProviderInterface_IsAutoAcceptEnabled_Call) Run(run func()) *ShipConnectionInfoProviderInterface_IsAutoAcceptEnabled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_IsAutoAcceptEnabled_Call) Return(_a0 bool) *ShipConnectionInfoProviderInterface_IsAutoAcceptEnabled_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_IsAutoAcceptEnabled_Call) RunAndReturn(run func() bool) *ShipConnectionInfoProviderInterface_IsAutoAcceptEnabled_Call {
	_c.Call.Return(run)
	return _c
}

// IsRemoteServiceForSKIPaired provides a mock function with given fields: _a0
func (_m *ShipConnectionInfoProviderInterface) IsRemoteServiceForSKIPaired(_a0 string) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for IsRemoteServiceForSKIPaired")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ShipConnectionInfoProviderInterface_IsRemoteServiceForSKIPaired_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRemoteServiceForSKIPaired'
type ShipConnectionInfoProviderInterface_IsRemoteServiceForSKIPaired_Call struct {
	*mock.Call
}

// IsRemoteServiceForSKIPaired is a helper method to define mock.On call
//   - _a0 string
func (_e *ShipConnectionInfoProviderInterface_Expecter) IsRemoteServiceForSKIPaired(_a0 interface{}) *ShipConnectionInfoProviderInterface_IsRemoteServiceForSKIPaired_Call {
	return &ShipConnectionInfoProviderInterface_IsRemoteServiceForSKIPaired_Call{Call: _e.mock.On("IsRemoteServiceForSKIPaired", _a0)}
}

func (_c *ShipConnectionInfoProviderInterface_IsRemoteServiceForSKIPaired_Call) Run(run func(_a0 string)) *ShipConnectionInfoProviderInterface_IsRemoteServiceForSKIPaired_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_IsRemoteServiceForSKIPaired_Call) Return(_a0 bool) *ShipConnectionInfoProviderInterface_IsRemoteServiceForSKIPaired_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_IsRemoteServiceForSKIPaired_Call) RunAndReturn(run func(string) bool) *ShipConnectionInfoProviderInterface_IsRemoteServiceForSKIPaired_Call {
	_c.Call.Return(run)
	return _c
}

// ReportServiceShipID provides a mock function with given fields: _a0, _a1
func (_m *ShipConnectionInfoProviderInterface) ReportServiceShipID(_a0 string, _a1 string) {
	_m.Called(_a0, _a1)
}

// ShipConnectionInfoProviderInterface_ReportServiceShipID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReportServiceShipID'
type ShipConnectionInfoProviderInterface_ReportServiceShipID_Call struct {
	*mock.Call
}

// ReportServiceShipID is a helper method to define mock.On call
//   - _a0 string
//   - _a1 string
func (_e *ShipConnectionInfoProviderInterface_Expecter) ReportServiceShipID(_a0 interface{}, _a1 interface{}) *ShipConnectionInfoProviderInterface_ReportServiceShipID_Call {
	return &ShipConnectionInfoProviderInterface_ReportServiceShipID_Call{Call: _e.mock.On("ReportServiceShipID", _a0, _a1)}
}

func (_c *ShipConnectionInfoProviderInterface_ReportServiceShipID_Call) Run(run func(_a0 string, _a1 string)) *ShipConnectionInfoProviderInterface_ReportServiceShipID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_ReportServiceShipID_Call) Return() *ShipConnectionInfoProviderInterface_ReportServiceShipID_Call {
	_c.Call.Return()
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_ReportServiceShipID_Call) RunAndReturn(run func(string, string)) *ShipConnectionInfoProviderInterface_ReportServiceShipID_Call {
	_c.Call.Return(run)
	return _c
}

// SetupRemoteDevice provides a mock function with given fields: ski, writeI
func (_m *ShipConnectionInfoProviderInterface) SetupRemoteDevice(ski string, writeI api.ShipConnectionDataWriterInterface) api.ShipConnectionDataReaderInterface {
	ret := _m.Called(ski, writeI)

	if len(ret) == 0 {
		panic("no return value specified for SetupRemoteDevice")
	}

	var r0 api.ShipConnectionDataReaderInterface
	if rf, ok := ret.Get(0).(func(string, api.ShipConnectionDataWriterInterface) api.ShipConnectionDataReaderInterface); ok {
		r0 = rf(ski, writeI)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.ShipConnectionDataReaderInterface)
		}
	}

	return r0
}

// ShipConnectionInfoProviderInterface_SetupRemoteDevice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetupRemoteDevice'
type ShipConnectionInfoProviderInterface_SetupRemoteDevice_Call struct {
	*mock.Call
}

// SetupRemoteDevice is a helper method to define mock.On call
//   - ski string
//   - writeI api.ShipConnectionDataWriterInterface
func (_e *ShipConnectionInfoProviderInterface_Expecter) SetupRemoteDevice(ski interface{}, writeI interface{}) *ShipConnectionInfoProviderInterface_SetupRemoteDevice_Call {
	return &ShipConnectionInfoProviderInterface_SetupRemoteDevice_Call{Call: _e.mock.On("SetupRemoteDevice", ski, writeI)}
}

func (_c *ShipConnectionInfoProviderInterface_SetupRemoteDevice_Call) Run(run func(ski string, writeI api.ShipConnectionDataWriterInterface)) *ShipConnectionInfoProviderInterface_SetupRemoteDevice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(api.ShipConnectionDataWriterInterface))
	})
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_SetupRemoteDevice_Call) Return(_a0 api.ShipConnectionDataReaderInterface) *ShipConnectionInfoProviderInterface_SetupRemoteDevice_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ShipConnectionInfoProviderInterface_SetupRemoteDevice_Call) RunAndReturn(run func(string, api.ShipConnectionDataWriterInterface) api.ShipConnectionDataReaderInterface) *ShipConnectionInfoProviderInterface_SetupRemoteDevice_Call {
	_c.Call.Return(run)
	return _c
}

// NewShipConnectionInfoProviderInterface creates a new instance of ShipConnectionInfoProviderInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewShipConnectionInfoProviderInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ShipConnectionInfoProviderInterface {
	mock := &ShipConnectionInfoProviderInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
