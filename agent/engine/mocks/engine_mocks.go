// Copyright 2015-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/engine (interfaces: TaskEngine,ImageManager)

// Package mock_engine is a generated GoMock package.
package mock_engine

import (
	context "context"
	reflect "reflect"

	api "github.com/aws/amazon-ecs-agent/agent/api"
	image "github.com/aws/amazon-ecs-agent/agent/engine/image"
	statechange "github.com/aws/amazon-ecs-agent/agent/statechange"
	statemanager "github.com/aws/amazon-ecs-agent/agent/statemanager"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskEngine is a mock of TaskEngine interface
type MockTaskEngine struct {
	ctrl     *gomock.Controller
	recorder *MockTaskEngineMockRecorder
}

// MockTaskEngineMockRecorder is the mock recorder for MockTaskEngine
type MockTaskEngineMockRecorder struct {
	mock *MockTaskEngine
}

// NewMockTaskEngine creates a new mock instance
func NewMockTaskEngine(ctrl *gomock.Controller) *MockTaskEngine {
	mock := &MockTaskEngine{ctrl: ctrl}
	mock.recorder = &MockTaskEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskEngine) EXPECT() *MockTaskEngineMockRecorder {
	return m.recorder
}

// AddTask mocks base method
func (m *MockTaskEngine) AddTask(arg0 *api.Task) error {
	ret := m.ctrl.Call(m, "AddTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTask indicates an expected call of AddTask
func (mr *MockTaskEngineMockRecorder) AddTask(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockTaskEngine)(nil).AddTask), arg0)
}

// Disable mocks base method
func (m *MockTaskEngine) Disable() {
	m.ctrl.Call(m, "Disable")
}

// Disable indicates an expected call of Disable
func (mr *MockTaskEngineMockRecorder) Disable() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockTaskEngine)(nil).Disable))
}

// GetTaskByArn mocks base method
func (m *MockTaskEngine) GetTaskByArn(arg0 string) (*api.Task, bool) {
	ret := m.ctrl.Call(m, "GetTaskByArn", arg0)
	ret0, _ := ret[0].(*api.Task)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetTaskByArn indicates an expected call of GetTaskByArn
func (mr *MockTaskEngineMockRecorder) GetTaskByArn(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByArn", reflect.TypeOf((*MockTaskEngine)(nil).GetTaskByArn), arg0)
}

// Init mocks base method
func (m *MockTaskEngine) Init(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockTaskEngineMockRecorder) Init(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockTaskEngine)(nil).Init), arg0)
}

// ListTasks mocks base method
func (m *MockTaskEngine) ListTasks() ([]*api.Task, error) {
	ret := m.ctrl.Call(m, "ListTasks")
	ret0, _ := ret[0].([]*api.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks
func (mr *MockTaskEngineMockRecorder) ListTasks() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockTaskEngine)(nil).ListTasks))
}

// MarshalJSON mocks base method
func (m *MockTaskEngine) MarshalJSON() ([]byte, error) {
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON
func (mr *MockTaskEngineMockRecorder) MarshalJSON() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockTaskEngine)(nil).MarshalJSON))
}

// MustInit mocks base method
func (m *MockTaskEngine) MustInit(arg0 context.Context) {
	m.ctrl.Call(m, "MustInit", arg0)
}

// MustInit indicates an expected call of MustInit
func (mr *MockTaskEngineMockRecorder) MustInit(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustInit", reflect.TypeOf((*MockTaskEngine)(nil).MustInit), arg0)
}

// SetSaver mocks base method
func (m *MockTaskEngine) SetSaver(arg0 statemanager.Saver) {
	m.ctrl.Call(m, "SetSaver", arg0)
}

// SetSaver indicates an expected call of SetSaver
func (mr *MockTaskEngineMockRecorder) SetSaver(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSaver", reflect.TypeOf((*MockTaskEngine)(nil).SetSaver), arg0)
}

// StateChangeEvents mocks base method
func (m *MockTaskEngine) StateChangeEvents() chan statechange.Event {
	ret := m.ctrl.Call(m, "StateChangeEvents")
	ret0, _ := ret[0].(chan statechange.Event)
	return ret0
}

// StateChangeEvents indicates an expected call of StateChangeEvents
func (mr *MockTaskEngineMockRecorder) StateChangeEvents() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateChangeEvents", reflect.TypeOf((*MockTaskEngine)(nil).StateChangeEvents))
}

// UnmarshalJSON mocks base method
func (m *MockTaskEngine) UnmarshalJSON(arg0 []byte) error {
	ret := m.ctrl.Call(m, "UnmarshalJSON", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmarshalJSON indicates an expected call of UnmarshalJSON
func (mr *MockTaskEngineMockRecorder) UnmarshalJSON(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalJSON", reflect.TypeOf((*MockTaskEngine)(nil).UnmarshalJSON), arg0)
}

// Version mocks base method
func (m *MockTaskEngine) Version() (string, error) {
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockTaskEngineMockRecorder) Version() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockTaskEngine)(nil).Version))
}

// MockImageManager is a mock of ImageManager interface
type MockImageManager struct {
	ctrl     *gomock.Controller
	recorder *MockImageManagerMockRecorder
}

// MockImageManagerMockRecorder is the mock recorder for MockImageManager
type MockImageManagerMockRecorder struct {
	mock *MockImageManager
}

// NewMockImageManager creates a new mock instance
func NewMockImageManager(ctrl *gomock.Controller) *MockImageManager {
	mock := &MockImageManager{ctrl: ctrl}
	mock.recorder = &MockImageManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageManager) EXPECT() *MockImageManagerMockRecorder {
	return m.recorder
}

// AddAllImageStates mocks base method
func (m *MockImageManager) AddAllImageStates(arg0 []*image.ImageState) {
	m.ctrl.Call(m, "AddAllImageStates", arg0)
}

// AddAllImageStates indicates an expected call of AddAllImageStates
func (mr *MockImageManagerMockRecorder) AddAllImageStates(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAllImageStates", reflect.TypeOf((*MockImageManager)(nil).AddAllImageStates), arg0)
}

// GetImageStateFromImageName mocks base method
func (m *MockImageManager) GetImageStateFromImageName(arg0 string) (*image.ImageState, bool) {
	ret := m.ctrl.Call(m, "GetImageStateFromImageName", arg0)
	ret0, _ := ret[0].(*image.ImageState)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetImageStateFromImageName indicates an expected call of GetImageStateFromImageName
func (mr *MockImageManagerMockRecorder) GetImageStateFromImageName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageStateFromImageName", reflect.TypeOf((*MockImageManager)(nil).GetImageStateFromImageName), arg0)
}

// RecordContainerReference mocks base method
func (m *MockImageManager) RecordContainerReference(arg0 *api.Container) error {
	ret := m.ctrl.Call(m, "RecordContainerReference", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordContainerReference indicates an expected call of RecordContainerReference
func (mr *MockImageManagerMockRecorder) RecordContainerReference(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordContainerReference", reflect.TypeOf((*MockImageManager)(nil).RecordContainerReference), arg0)
}

// RemoveContainerReferenceFromImageState mocks base method
func (m *MockImageManager) RemoveContainerReferenceFromImageState(arg0 *api.Container) error {
	ret := m.ctrl.Call(m, "RemoveContainerReferenceFromImageState", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainerReferenceFromImageState indicates an expected call of RemoveContainerReferenceFromImageState
func (mr *MockImageManagerMockRecorder) RemoveContainerReferenceFromImageState(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainerReferenceFromImageState", reflect.TypeOf((*MockImageManager)(nil).RemoveContainerReferenceFromImageState), arg0)
}

// SetSaver mocks base method
func (m *MockImageManager) SetSaver(arg0 statemanager.Saver) {
	m.ctrl.Call(m, "SetSaver", arg0)
}

// SetSaver indicates an expected call of SetSaver
func (mr *MockImageManagerMockRecorder) SetSaver(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSaver", reflect.TypeOf((*MockImageManager)(nil).SetSaver), arg0)
}

// StartImageCleanupProcess mocks base method
func (m *MockImageManager) StartImageCleanupProcess(arg0 context.Context) {
	m.ctrl.Call(m, "StartImageCleanupProcess", arg0)
}

// StartImageCleanupProcess indicates an expected call of StartImageCleanupProcess
func (mr *MockImageManagerMockRecorder) StartImageCleanupProcess(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartImageCleanupProcess", reflect.TypeOf((*MockImageManager)(nil).StartImageCleanupProcess), arg0)
}
