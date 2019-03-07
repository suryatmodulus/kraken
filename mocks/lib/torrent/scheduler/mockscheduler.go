// Copyright (c) 2016-2019 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/kraken/lib/torrent/scheduler (interfaces: ReloadableScheduler,Scheduler)

// Package mockscheduler is a generated GoMock package.
package mockscheduler

import (
	gomock "github.com/golang/mock/gomock"
	core "github.com/uber/kraken/core"
	scheduler "github.com/uber/kraken/lib/torrent/scheduler"
	connstate "github.com/uber/kraken/lib/torrent/scheduler/connstate"
	reflect "reflect"
)

// MockReloadableScheduler is a mock of ReloadableScheduler interface
type MockReloadableScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockReloadableSchedulerMockRecorder
}

// MockReloadableSchedulerMockRecorder is the mock recorder for MockReloadableScheduler
type MockReloadableSchedulerMockRecorder struct {
	mock *MockReloadableScheduler
}

// NewMockReloadableScheduler creates a new mock instance
func NewMockReloadableScheduler(ctrl *gomock.Controller) *MockReloadableScheduler {
	mock := &MockReloadableScheduler{ctrl: ctrl}
	mock.recorder = &MockReloadableSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReloadableScheduler) EXPECT() *MockReloadableSchedulerMockRecorder {
	return m.recorder
}

// BlacklistSnapshot mocks base method
func (m *MockReloadableScheduler) BlacklistSnapshot() ([]connstate.BlacklistedConn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlacklistSnapshot")
	ret0, _ := ret[0].([]connstate.BlacklistedConn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlacklistSnapshot indicates an expected call of BlacklistSnapshot
func (mr *MockReloadableSchedulerMockRecorder) BlacklistSnapshot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlacklistSnapshot", reflect.TypeOf((*MockReloadableScheduler)(nil).BlacklistSnapshot))
}

// Download mocks base method
func (m *MockReloadableScheduler) Download(arg0 string, arg1 core.Digest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download
func (mr *MockReloadableSchedulerMockRecorder) Download(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockReloadableScheduler)(nil).Download), arg0, arg1)
}

// Probe mocks base method
func (m *MockReloadableScheduler) Probe() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Probe")
	ret0, _ := ret[0].(error)
	return ret0
}

// Probe indicates an expected call of Probe
func (mr *MockReloadableSchedulerMockRecorder) Probe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Probe", reflect.TypeOf((*MockReloadableScheduler)(nil).Probe))
}

// Reload mocks base method
func (m *MockReloadableScheduler) Reload(arg0 scheduler.Config) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reload", arg0)
}

// Reload indicates an expected call of Reload
func (mr *MockReloadableSchedulerMockRecorder) Reload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockReloadableScheduler)(nil).Reload), arg0)
}

// RemoveTorrent mocks base method
func (m *MockReloadableScheduler) RemoveTorrent(arg0 core.Digest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTorrent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTorrent indicates an expected call of RemoveTorrent
func (mr *MockReloadableSchedulerMockRecorder) RemoveTorrent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTorrent", reflect.TypeOf((*MockReloadableScheduler)(nil).RemoveTorrent), arg0)
}

// Stop mocks base method
func (m *MockReloadableScheduler) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockReloadableSchedulerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockReloadableScheduler)(nil).Stop))
}

// MockScheduler is a mock of Scheduler interface
type MockScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerMockRecorder
}

// MockSchedulerMockRecorder is the mock recorder for MockScheduler
type MockSchedulerMockRecorder struct {
	mock *MockScheduler
}

// NewMockScheduler creates a new mock instance
func NewMockScheduler(ctrl *gomock.Controller) *MockScheduler {
	mock := &MockScheduler{ctrl: ctrl}
	mock.recorder = &MockSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScheduler) EXPECT() *MockSchedulerMockRecorder {
	return m.recorder
}

// BlacklistSnapshot mocks base method
func (m *MockScheduler) BlacklistSnapshot() ([]connstate.BlacklistedConn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlacklistSnapshot")
	ret0, _ := ret[0].([]connstate.BlacklistedConn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlacklistSnapshot indicates an expected call of BlacklistSnapshot
func (mr *MockSchedulerMockRecorder) BlacklistSnapshot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlacklistSnapshot", reflect.TypeOf((*MockScheduler)(nil).BlacklistSnapshot))
}

// Download mocks base method
func (m *MockScheduler) Download(arg0 string, arg1 core.Digest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download
func (mr *MockSchedulerMockRecorder) Download(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockScheduler)(nil).Download), arg0, arg1)
}

// Probe mocks base method
func (m *MockScheduler) Probe() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Probe")
	ret0, _ := ret[0].(error)
	return ret0
}

// Probe indicates an expected call of Probe
func (mr *MockSchedulerMockRecorder) Probe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Probe", reflect.TypeOf((*MockScheduler)(nil).Probe))
}

// RemoveTorrent mocks base method
func (m *MockScheduler) RemoveTorrent(arg0 core.Digest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTorrent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTorrent indicates an expected call of RemoveTorrent
func (mr *MockSchedulerMockRecorder) RemoveTorrent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTorrent", reflect.TypeOf((*MockScheduler)(nil).RemoveTorrent), arg0)
}

// Stop mocks base method
func (m *MockScheduler) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockSchedulerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockScheduler)(nil).Stop))
}
