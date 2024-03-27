// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ShipConnectionDataWriterInterface is an autogenerated mock type for the ShipConnectionDataWriterInterface type
type ShipConnectionDataWriterInterface struct {
	mock.Mock
}

type ShipConnectionDataWriterInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ShipConnectionDataWriterInterface) EXPECT() *ShipConnectionDataWriterInterface_Expecter {
	return &ShipConnectionDataWriterInterface_Expecter{mock: &_m.Mock}
}

// WriteShipMessageWithPayload provides a mock function with given fields: message
func (_m *ShipConnectionDataWriterInterface) WriteShipMessageWithPayload(message []byte) {
	_m.Called(message)
}

// ShipConnectionDataWriterInterface_WriteShipMessageWithPayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteShipMessageWithPayload'
type ShipConnectionDataWriterInterface_WriteShipMessageWithPayload_Call struct {
	*mock.Call
}

// WriteShipMessageWithPayload is a helper method to define mock.On call
//   - message []byte
func (_e *ShipConnectionDataWriterInterface_Expecter) WriteShipMessageWithPayload(message interface{}) *ShipConnectionDataWriterInterface_WriteShipMessageWithPayload_Call {
	return &ShipConnectionDataWriterInterface_WriteShipMessageWithPayload_Call{Call: _e.mock.On("WriteShipMessageWithPayload", message)}
}

func (_c *ShipConnectionDataWriterInterface_WriteShipMessageWithPayload_Call) Run(run func(message []byte)) *ShipConnectionDataWriterInterface_WriteShipMessageWithPayload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *ShipConnectionDataWriterInterface_WriteShipMessageWithPayload_Call) Return() *ShipConnectionDataWriterInterface_WriteShipMessageWithPayload_Call {
	_c.Call.Return()
	return _c
}

func (_c *ShipConnectionDataWriterInterface_WriteShipMessageWithPayload_Call) RunAndReturn(run func([]byte)) *ShipConnectionDataWriterInterface_WriteShipMessageWithPayload_Call {
	_c.Call.Return(run)
	return _c
}

// NewShipConnectionDataWriterInterface creates a new instance of ShipConnectionDataWriterInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewShipConnectionDataWriterInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ShipConnectionDataWriterInterface {
	mock := &ShipConnectionDataWriterInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
