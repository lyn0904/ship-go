// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/ship-go/api"
	mock "github.com/stretchr/testify/mock"
)

// MdnsInterface is an autogenerated mock type for the MdnsInterface type
type MdnsInterface struct {
	mock.Mock
}

type MdnsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MdnsInterface) EXPECT() *MdnsInterface_Expecter {
	return &MdnsInterface_Expecter{mock: &_m.Mock}
}

// AnnounceMdnsEntry provides a mock function with given fields:
func (_m *MdnsInterface) AnnounceMdnsEntry() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AnnounceMdnsEntry")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MdnsInterface_AnnounceMdnsEntry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AnnounceMdnsEntry'
type MdnsInterface_AnnounceMdnsEntry_Call struct {
	*mock.Call
}

// AnnounceMdnsEntry is a helper method to define mock.On call
func (_e *MdnsInterface_Expecter) AnnounceMdnsEntry() *MdnsInterface_AnnounceMdnsEntry_Call {
	return &MdnsInterface_AnnounceMdnsEntry_Call{Call: _e.mock.On("AnnounceMdnsEntry")}
}

func (_c *MdnsInterface_AnnounceMdnsEntry_Call) Run(run func()) *MdnsInterface_AnnounceMdnsEntry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MdnsInterface_AnnounceMdnsEntry_Call) Return(_a0 error) *MdnsInterface_AnnounceMdnsEntry_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MdnsInterface_AnnounceMdnsEntry_Call) RunAndReturn(run func() error) *MdnsInterface_AnnounceMdnsEntry_Call {
	_c.Call.Return(run)
	return _c
}

// RequestMdnsEntries provides a mock function with given fields:
func (_m *MdnsInterface) RequestMdnsEntries() {
	_m.Called()
}

// MdnsInterface_RequestMdnsEntries_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestMdnsEntries'
type MdnsInterface_RequestMdnsEntries_Call struct {
	*mock.Call
}

// RequestMdnsEntries is a helper method to define mock.On call
func (_e *MdnsInterface_Expecter) RequestMdnsEntries() *MdnsInterface_RequestMdnsEntries_Call {
	return &MdnsInterface_RequestMdnsEntries_Call{Call: _e.mock.On("RequestMdnsEntries")}
}

func (_c *MdnsInterface_RequestMdnsEntries_Call) Run(run func()) *MdnsInterface_RequestMdnsEntries_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MdnsInterface_RequestMdnsEntries_Call) Return() *MdnsInterface_RequestMdnsEntries_Call {
	_c.Call.Return()
	return _c
}

func (_c *MdnsInterface_RequestMdnsEntries_Call) RunAndReturn(run func()) *MdnsInterface_RequestMdnsEntries_Call {
	_c.Call.Return(run)
	return _c
}

// SetAutoAccept provides a mock function with given fields: _a0
func (_m *MdnsInterface) SetAutoAccept(_a0 bool) {
	_m.Called(_a0)
}

// MdnsInterface_SetAutoAccept_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAutoAccept'
type MdnsInterface_SetAutoAccept_Call struct {
	*mock.Call
}

// SetAutoAccept is a helper method to define mock.On call
//   - _a0 bool
func (_e *MdnsInterface_Expecter) SetAutoAccept(_a0 interface{}) *MdnsInterface_SetAutoAccept_Call {
	return &MdnsInterface_SetAutoAccept_Call{Call: _e.mock.On("SetAutoAccept", _a0)}
}

func (_c *MdnsInterface_SetAutoAccept_Call) Run(run func(_a0 bool)) *MdnsInterface_SetAutoAccept_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MdnsInterface_SetAutoAccept_Call) Return() *MdnsInterface_SetAutoAccept_Call {
	_c.Call.Return()
	return _c
}

func (_c *MdnsInterface_SetAutoAccept_Call) RunAndReturn(run func(bool)) *MdnsInterface_SetAutoAccept_Call {
	_c.Call.Return(run)
	return _c
}

// Shutdown provides a mock function with given fields:
func (_m *MdnsInterface) Shutdown() {
	_m.Called()
}

// MdnsInterface_Shutdown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Shutdown'
type MdnsInterface_Shutdown_Call struct {
	*mock.Call
}

// Shutdown is a helper method to define mock.On call
func (_e *MdnsInterface_Expecter) Shutdown() *MdnsInterface_Shutdown_Call {
	return &MdnsInterface_Shutdown_Call{Call: _e.mock.On("Shutdown")}
}

func (_c *MdnsInterface_Shutdown_Call) Run(run func()) *MdnsInterface_Shutdown_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MdnsInterface_Shutdown_Call) Return() *MdnsInterface_Shutdown_Call {
	_c.Call.Return()
	return _c
}

func (_c *MdnsInterface_Shutdown_Call) RunAndReturn(run func()) *MdnsInterface_Shutdown_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: cb
func (_m *MdnsInterface) Start(cb api.MdnsReportInterface) error {
	ret := _m.Called(cb)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(api.MdnsReportInterface) error); ok {
		r0 = rf(cb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MdnsInterface_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MdnsInterface_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - cb api.MdnsReportInterface
func (_e *MdnsInterface_Expecter) Start(cb interface{}) *MdnsInterface_Start_Call {
	return &MdnsInterface_Start_Call{Call: _e.mock.On("Start", cb)}
}

func (_c *MdnsInterface_Start_Call) Run(run func(cb api.MdnsReportInterface)) *MdnsInterface_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.MdnsReportInterface))
	})
	return _c
}

func (_c *MdnsInterface_Start_Call) Return(_a0 error) *MdnsInterface_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MdnsInterface_Start_Call) RunAndReturn(run func(api.MdnsReportInterface) error) *MdnsInterface_Start_Call {
	_c.Call.Return(run)
	return _c
}

// UnannounceMdnsEntry provides a mock function with given fields:
func (_m *MdnsInterface) UnannounceMdnsEntry() {
	_m.Called()
}

// MdnsInterface_UnannounceMdnsEntry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnannounceMdnsEntry'
type MdnsInterface_UnannounceMdnsEntry_Call struct {
	*mock.Call
}

// UnannounceMdnsEntry is a helper method to define mock.On call
func (_e *MdnsInterface_Expecter) UnannounceMdnsEntry() *MdnsInterface_UnannounceMdnsEntry_Call {
	return &MdnsInterface_UnannounceMdnsEntry_Call{Call: _e.mock.On("UnannounceMdnsEntry")}
}

func (_c *MdnsInterface_UnannounceMdnsEntry_Call) Run(run func()) *MdnsInterface_UnannounceMdnsEntry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MdnsInterface_UnannounceMdnsEntry_Call) Return() *MdnsInterface_UnannounceMdnsEntry_Call {
	_c.Call.Return()
	return _c
}

func (_c *MdnsInterface_UnannounceMdnsEntry_Call) RunAndReturn(run func()) *MdnsInterface_UnannounceMdnsEntry_Call {
	_c.Call.Return(run)
	return _c
}

// NewMdnsInterface creates a new instance of MdnsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMdnsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MdnsInterface {
	mock := &MdnsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
