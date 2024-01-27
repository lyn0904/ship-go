// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/enbility/ship-go/api (interfaces: MdnsInterface,HubReaderInterface)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/mockgen_api.go -package=mocks github.com/enbility/ship-go/api MdnsInterface,HubReaderInterface
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	api "github.com/enbility/ship-go/api"
	gomock "go.uber.org/mock/gomock"
)

// MockMdnsInterface is a mock of MdnsInterface interface.
type MockMdnsInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMdnsInterfaceMockRecorder
}

// MockMdnsInterfaceMockRecorder is the mock recorder for MockMdnsInterface.
type MockMdnsInterfaceMockRecorder struct {
	mock *MockMdnsInterface
}

// NewMockMdnsInterface creates a new mock instance.
func NewMockMdnsInterface(ctrl *gomock.Controller) *MockMdnsInterface {
	mock := &MockMdnsInterface{ctrl: ctrl}
	mock.recorder = &MockMdnsInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMdnsInterface) EXPECT() *MockMdnsInterfaceMockRecorder {
	return m.recorder
}

// AnnounceMdnsEntry mocks base method.
func (m *MockMdnsInterface) AnnounceMdnsEntry() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnnounceMdnsEntry")
	ret0, _ := ret[0].(error)
	return ret0
}

// AnnounceMdnsEntry indicates an expected call of AnnounceMdnsEntry.
func (mr *MockMdnsInterfaceMockRecorder) AnnounceMdnsEntry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceMdnsEntry", reflect.TypeOf((*MockMdnsInterface)(nil).AnnounceMdnsEntry))
}

// RequestMdnsEntries mocks base method.
func (m *MockMdnsInterface) RequestMdnsEntries() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RequestMdnsEntries")
}

// RequestMdnsEntries indicates an expected call of RequestMdnsEntries.
func (mr *MockMdnsInterfaceMockRecorder) RequestMdnsEntries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestMdnsEntries", reflect.TypeOf((*MockMdnsInterface)(nil).RequestMdnsEntries))
}

// SetAutoAccept mocks base method.
func (m *MockMdnsInterface) SetAutoAccept(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAutoAccept", arg0)
}

// SetAutoAccept indicates an expected call of SetAutoAccept.
func (mr *MockMdnsInterfaceMockRecorder) SetAutoAccept(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAutoAccept", reflect.TypeOf((*MockMdnsInterface)(nil).SetAutoAccept), arg0)
}

// Shutdown mocks base method.
func (m *MockMdnsInterface) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockMdnsInterfaceMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockMdnsInterface)(nil).Shutdown))
}

// Start mocks base method.
func (m *MockMdnsInterface) Start(arg0 api.MdnsReportInterface) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockMdnsInterfaceMockRecorder) Start(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockMdnsInterface)(nil).Start), arg0)
}

// UnannounceMdnsEntry mocks base method.
func (m *MockMdnsInterface) UnannounceMdnsEntry() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnannounceMdnsEntry")
}

// UnannounceMdnsEntry indicates an expected call of UnannounceMdnsEntry.
func (mr *MockMdnsInterfaceMockRecorder) UnannounceMdnsEntry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnannounceMdnsEntry", reflect.TypeOf((*MockMdnsInterface)(nil).UnannounceMdnsEntry))
}

// MockHubReaderInterface is a mock of HubReaderInterface interface.
type MockHubReaderInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHubReaderInterfaceMockRecorder
}

// MockHubReaderInterfaceMockRecorder is the mock recorder for MockHubReaderInterface.
type MockHubReaderInterfaceMockRecorder struct {
	mock *MockHubReaderInterface
}

// NewMockHubReaderInterface creates a new mock instance.
func NewMockHubReaderInterface(ctrl *gomock.Controller) *MockHubReaderInterface {
	mock := &MockHubReaderInterface{ctrl: ctrl}
	mock.recorder = &MockHubReaderInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHubReaderInterface) EXPECT() *MockHubReaderInterfaceMockRecorder {
	return m.recorder
}

// AllowWaitingForTrust mocks base method.
func (m *MockHubReaderInterface) AllowWaitingForTrust(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllowWaitingForTrust", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AllowWaitingForTrust indicates an expected call of AllowWaitingForTrust.
func (mr *MockHubReaderInterfaceMockRecorder) AllowWaitingForTrust(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowWaitingForTrust", reflect.TypeOf((*MockHubReaderInterface)(nil).AllowWaitingForTrust), arg0)
}

// RemoteSKIConnected mocks base method.
func (m *MockHubReaderInterface) RemoteSKIConnected(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoteSKIConnected", arg0)
}

// RemoteSKIConnected indicates an expected call of RemoteSKIConnected.
func (mr *MockHubReaderInterfaceMockRecorder) RemoteSKIConnected(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteSKIConnected", reflect.TypeOf((*MockHubReaderInterface)(nil).RemoteSKIConnected), arg0)
}

// RemoteSKIDisconnected mocks base method.
func (m *MockHubReaderInterface) RemoteSKIDisconnected(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoteSKIDisconnected", arg0)
}

// RemoteSKIDisconnected indicates an expected call of RemoteSKIDisconnected.
func (mr *MockHubReaderInterfaceMockRecorder) RemoteSKIDisconnected(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteSKIDisconnected", reflect.TypeOf((*MockHubReaderInterface)(nil).RemoteSKIDisconnected), arg0)
}

// ServicePairingDetailUpdate mocks base method.
func (m *MockHubReaderInterface) ServicePairingDetailUpdate(arg0 string, arg1 *api.ConnectionStateDetail) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServicePairingDetailUpdate", arg0, arg1)
}

// ServicePairingDetailUpdate indicates an expected call of ServicePairingDetailUpdate.
func (mr *MockHubReaderInterfaceMockRecorder) ServicePairingDetailUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServicePairingDetailUpdate", reflect.TypeOf((*MockHubReaderInterface)(nil).ServicePairingDetailUpdate), arg0, arg1)
}

// ServiceShipIDUpdate mocks base method.
func (m *MockHubReaderInterface) ServiceShipIDUpdate(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServiceShipIDUpdate", arg0, arg1)
}

// ServiceShipIDUpdate indicates an expected call of ServiceShipIDUpdate.
func (mr *MockHubReaderInterfaceMockRecorder) ServiceShipIDUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceShipIDUpdate", reflect.TypeOf((*MockHubReaderInterface)(nil).ServiceShipIDUpdate), arg0, arg1)
}

// SetupRemoteDevice mocks base method.
func (m *MockHubReaderInterface) SetupRemoteDevice(arg0 string, arg1 api.ShipConnectionDataWriterInterface) api.ShipConnectionDataReaderInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupRemoteDevice", arg0, arg1)
	ret0, _ := ret[0].(api.ShipConnectionDataReaderInterface)
	return ret0
}

// SetupRemoteDevice indicates an expected call of SetupRemoteDevice.
func (mr *MockHubReaderInterfaceMockRecorder) SetupRemoteDevice(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupRemoteDevice", reflect.TypeOf((*MockHubReaderInterface)(nil).SetupRemoteDevice), arg0, arg1)
}

// VisibleMDNSRecordsUpdated mocks base method.
func (m *MockHubReaderInterface) VisibleMDNSRecordsUpdated(arg0 []*api.MdnsEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VisibleMDNSRecordsUpdated", arg0)
}

// VisibleMDNSRecordsUpdated indicates an expected call of VisibleMDNSRecordsUpdated.
func (mr *MockHubReaderInterfaceMockRecorder) VisibleMDNSRecordsUpdated(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VisibleMDNSRecordsUpdated", reflect.TypeOf((*MockHubReaderInterface)(nil).VisibleMDNSRecordsUpdated), arg0)
}
