// Code generated by MockGen. DO NOT EDIT.
// Source: consul.go

package consul

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/hashicorp/consul/api"
	reflect "reflect"
)

// mockConsulClient is a mock of consulClient interface
type mockConsulClient struct {
	ctrl     *gomock.Controller
	recorder *mockConsulClientMockRecorder
}

// mockConsulClientMockRecorder is the mock recorder for mockConsulClient
type mockConsulClientMockRecorder struct {
	mock *mockConsulClient
}

// newMockConsulClient creates a new mock instance
func newMockConsulClient(ctrl *gomock.Controller) *mockConsulClient {
	mock := &mockConsulClient{ctrl: ctrl}
	mock.recorder = &mockConsulClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *mockConsulClient) EXPECT() *mockConsulClientMockRecorder {
	return m.recorder
}

// Catalog mocks base method
func (m *mockConsulClient) Catalog() catalogInterface {
	ret := m.ctrl.Call(m, "Catalog")
	ret0, _ := ret[0].(catalogInterface)
	return ret0
}

// Catalog indicates an expected call of Catalog
func (mr *mockConsulClientMockRecorder) Catalog() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Catalog", reflect.TypeOf((*mockConsulClient)(nil).Catalog))
}

// Health mocks base method
func (m *mockConsulClient) Health() healthInterface {
	ret := m.ctrl.Call(m, "Health")
	ret0, _ := ret[0].(healthInterface)
	return ret0
}

// Health indicates an expected call of Health
func (mr *mockConsulClientMockRecorder) Health() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*mockConsulClient)(nil).Health))
}

// mockCatalogInterface is a mock of catalogInterface interface
type mockCatalogInterface struct {
	ctrl     *gomock.Controller
	recorder *mockCatalogInterfaceMockRecorder
}

// mockCatalogInterfaceMockRecorder is the mock recorder for mockCatalogInterface
type mockCatalogInterfaceMockRecorder struct {
	mock *mockCatalogInterface
}

// newMockCatalogInterface creates a new mock instance
func newMockCatalogInterface(ctrl *gomock.Controller) *mockCatalogInterface {
	mock := &mockCatalogInterface{ctrl: ctrl}
	mock.recorder = &mockCatalogInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *mockCatalogInterface) EXPECT() *mockCatalogInterfaceMockRecorder {
	return m.recorder
}

// Datacenters mocks base method
func (m *mockCatalogInterface) Datacenters() ([]string, error) {
	ret := m.ctrl.Call(m, "Datacenters")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Datacenters indicates an expected call of Datacenters
func (mr *mockCatalogInterfaceMockRecorder) Datacenters() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Datacenters", reflect.TypeOf((*mockCatalogInterface)(nil).Datacenters))
}

// Services mocks base method
func (m *mockCatalogInterface) Services(queryOpts *api.QueryOptions) (map[string][]string, *api.QueryMeta, error) {
	ret := m.ctrl.Call(m, "Services", queryOpts)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(*api.QueryMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Services indicates an expected call of Services
func (mr *mockCatalogInterfaceMockRecorder) Services(queryOpts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Services", reflect.TypeOf((*mockCatalogInterface)(nil).Services), queryOpts)
}

// Service mocks base method
func (m *mockCatalogInterface) Service(service, tag string, queryOpts *api.QueryOptions) ([]*api.CatalogService, *api.QueryMeta, error) {
	ret := m.ctrl.Call(m, "Service", service, tag, queryOpts)
	ret0, _ := ret[0].([]*api.CatalogService)
	ret1, _ := ret[1].(*api.QueryMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Service indicates an expected call of Service
func (mr *mockCatalogInterfaceMockRecorder) Service(service, tag, queryOpts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Service", reflect.TypeOf((*mockCatalogInterface)(nil).Service), service, tag, queryOpts)
}

// mockHealthInterface is a mock of healthInterface interface
type mockHealthInterface struct {
	ctrl     *gomock.Controller
	recorder *mockHealthInterfaceMockRecorder
}

// mockHealthInterfaceMockRecorder is the mock recorder for mockHealthInterface
type mockHealthInterfaceMockRecorder struct {
	mock *mockHealthInterface
}

// newMockHealthInterface creates a new mock instance
func newMockHealthInterface(ctrl *gomock.Controller) *mockHealthInterface {
	mock := &mockHealthInterface{ctrl: ctrl}
	mock.recorder = &mockHealthInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *mockHealthInterface) EXPECT() *mockHealthInterfaceMockRecorder {
	return m.recorder
}

// Node mocks base method
func (m *mockHealthInterface) Node(node string, queryOpts *api.QueryOptions) (api.HealthChecks, *api.QueryMeta, error) {
	ret := m.ctrl.Call(m, "Node", node, queryOpts)
	ret0, _ := ret[0].(api.HealthChecks)
	ret1, _ := ret[1].(*api.QueryMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Node indicates an expected call of Node
func (mr *mockHealthInterfaceMockRecorder) Node(node, queryOpts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*mockHealthInterface)(nil).Node), node, queryOpts)
}

// mockGetClientInterface is a mock of getClientInterface interface
type mockGetClientInterface struct {
	ctrl     *gomock.Controller
	recorder *mockGetClientInterfaceMockRecorder
}

// mockGetClientInterfaceMockRecorder is the mock recorder for mockGetClientInterface
type mockGetClientInterfaceMockRecorder struct {
	mock *mockGetClientInterface
}

// newMockGetClientInterface creates a new mock instance
func newMockGetClientInterface(ctrl *gomock.Controller) *mockGetClientInterface {
	mock := &mockGetClientInterface{ctrl: ctrl}
	mock.recorder = &mockGetClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *mockGetClientInterface) EXPECT() *mockGetClientInterfaceMockRecorder {
	return m.recorder
}

// getClient mocks base method
func (m *mockGetClientInterface) getClient() (consulClient, error) {
	ret := m.ctrl.Call(m, "getClient")
	ret0, _ := ret[0].(consulClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getClient indicates an expected call of getClient
func (mr *mockGetClientInterfaceMockRecorder) getClient() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getClient", reflect.TypeOf((*mockGetClientInterface)(nil).getClient))
}
