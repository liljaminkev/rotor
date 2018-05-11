// Code generated by MockGen. DO NOT EDIT.
// Source: snapshot_cache.go

package adapter

import (
	cache "github.com/envoyproxy/go-control-plane/pkg/cache"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// mockSnapshotCache is a mock of snapshotCache interface
type mockSnapshotCache struct {
	ctrl     *gomock.Controller
	recorder *mockSnapshotCacheMockRecorder
}

// mockSnapshotCacheMockRecorder is the mock recorder for mockSnapshotCache
type mockSnapshotCacheMockRecorder struct {
	mock *mockSnapshotCache
}

// newMockSnapshotCache creates a new mock instance
func newMockSnapshotCache(ctrl *gomock.Controller) *mockSnapshotCache {
	mock := &mockSnapshotCache{ctrl: ctrl}
	mock.recorder = &mockSnapshotCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *mockSnapshotCache) EXPECT() *mockSnapshotCacheMockRecorder {
	return m.recorder
}

// SetSnapshot mocks base method
func (m *mockSnapshotCache) SetSnapshot(node string, snapshot cache.Snapshot) error {
	ret := m.ctrl.Call(m, "SetSnapshot", node, snapshot)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSnapshot indicates an expected call of SetSnapshot
func (mr *mockSnapshotCacheMockRecorder) SetSnapshot(node, snapshot interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSnapshot", reflect.TypeOf((*mockSnapshotCache)(nil).SetSnapshot), node, snapshot)
}

// ClearSnapshot mocks base method
func (m *mockSnapshotCache) ClearSnapshot(node string) {
	m.ctrl.Call(m, "ClearSnapshot", node)
}

// ClearSnapshot indicates an expected call of ClearSnapshot
func (mr *mockSnapshotCacheMockRecorder) ClearSnapshot(node interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearSnapshot", reflect.TypeOf((*mockSnapshotCache)(nil).ClearSnapshot), node)
}