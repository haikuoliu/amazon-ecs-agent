// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
// Source: github.com/aws/amazon-ecs-agent/agent/dockerclient/dockerapi (interfaces: DockerClient)

// Package mock_dockerapi is a generated GoMock package.
package mock_dockerapi

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	container "github.com/aws/amazon-ecs-agent/agent/api/container"
	status "github.com/aws/amazon-ecs-agent/agent/api/container/status"
	dockerclient "github.com/aws/amazon-ecs-agent/agent/dockerclient"
	dockerapi "github.com/aws/amazon-ecs-agent/agent/dockerclient/dockerapi"
	types "github.com/docker/docker/api/types"
	container0 "github.com/docker/docker/api/types/container"
	filters "github.com/docker/docker/api/types/filters"
	gomock "github.com/golang/mock/gomock"
)

// MockDockerClient is a mock of DockerClient interface
type MockDockerClient struct {
	ctrl     *gomock.Controller
	recorder *MockDockerClientMockRecorder
}

// MockDockerClientMockRecorder is the mock recorder for MockDockerClient
type MockDockerClientMockRecorder struct {
	mock *MockDockerClient
}

// NewMockDockerClient creates a new mock instance
func NewMockDockerClient(ctrl *gomock.Controller) *MockDockerClient {
	mock := &MockDockerClient{ctrl: ctrl}
	mock.recorder = &MockDockerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDockerClient) EXPECT() *MockDockerClientMockRecorder {
	return m.recorder
}

// APIVersion mocks base method
func (m *MockDockerClient) APIVersion() (dockerclient.DockerVersion, error) {
	ret := m.ctrl.Call(m, "APIVersion")
	ret0, _ := ret[0].(dockerclient.DockerVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIVersion indicates an expected call of APIVersion
func (mr *MockDockerClientMockRecorder) APIVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIVersion", reflect.TypeOf((*MockDockerClient)(nil).APIVersion))
}

// ContainerEvents mocks base method
func (m *MockDockerClient) ContainerEvents(arg0 context.Context) (<-chan dockerapi.DockerContainerChangeEvent, error) {
	ret := m.ctrl.Call(m, "ContainerEvents", arg0)
	ret0, _ := ret[0].(<-chan dockerapi.DockerContainerChangeEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerEvents indicates an expected call of ContainerEvents
func (mr *MockDockerClientMockRecorder) ContainerEvents(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerEvents", reflect.TypeOf((*MockDockerClient)(nil).ContainerEvents), arg0)
}

// CreateContainer mocks base method
func (m *MockDockerClient) CreateContainer(arg0 context.Context, arg1 *container0.Config, arg2 *container0.HostConfig, arg3 string, arg4 time.Duration) dockerapi.DockerContainerMetadata {
	ret := m.ctrl.Call(m, "CreateContainer", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(dockerapi.DockerContainerMetadata)
	return ret0
}

// CreateContainer indicates an expected call of CreateContainer
func (mr *MockDockerClientMockRecorder) CreateContainer(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockDockerClient)(nil).CreateContainer), arg0, arg1, arg2, arg3, arg4)
}

// CreateVolume mocks base method
func (m *MockDockerClient) CreateVolume(arg0 context.Context, arg1, arg2 string, arg3, arg4 map[string]string, arg5 time.Duration) dockerapi.SDKVolumeResponse {
	ret := m.ctrl.Call(m, "CreateVolume", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(dockerapi.SDKVolumeResponse)
	return ret0
}

// CreateVolume indicates an expected call of CreateVolume
func (mr *MockDockerClientMockRecorder) CreateVolume(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockDockerClient)(nil).CreateVolume), arg0, arg1, arg2, arg3, arg4, arg5)
}

// DescribeContainer mocks base method
func (m *MockDockerClient) DescribeContainer(arg0 context.Context, arg1 string) (status.ContainerStatus, dockerapi.DockerContainerMetadata) {
	ret := m.ctrl.Call(m, "DescribeContainer", arg0, arg1)
	ret0, _ := ret[0].(status.ContainerStatus)
	ret1, _ := ret[1].(dockerapi.DockerContainerMetadata)
	return ret0, ret1
}

// DescribeContainer indicates an expected call of DescribeContainer
func (mr *MockDockerClientMockRecorder) DescribeContainer(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeContainer", reflect.TypeOf((*MockDockerClient)(nil).DescribeContainer), arg0, arg1)
}

// InspectContainer mocks base method
func (m *MockDockerClient) InspectContainer(arg0 context.Context, arg1 string, arg2 time.Duration) (*types.ContainerJSON, error) {
	ret := m.ctrl.Call(m, "InspectContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.ContainerJSON)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectContainer indicates an expected call of InspectContainer
func (mr *MockDockerClientMockRecorder) InspectContainer(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectContainer", reflect.TypeOf((*MockDockerClient)(nil).InspectContainer), arg0, arg1, arg2)
}

// InspectImage mocks base method
func (m *MockDockerClient) InspectImage(arg0 string) (*types.ImageInspect, error) {
	ret := m.ctrl.Call(m, "InspectImage", arg0)
	ret0, _ := ret[0].(*types.ImageInspect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectImage indicates an expected call of InspectImage
func (mr *MockDockerClientMockRecorder) InspectImage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectImage", reflect.TypeOf((*MockDockerClient)(nil).InspectImage), arg0)
}

// InspectVolume mocks base method
func (m *MockDockerClient) InspectVolume(arg0 context.Context, arg1 string, arg2 time.Duration) dockerapi.SDKVolumeResponse {
	ret := m.ctrl.Call(m, "InspectVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(dockerapi.SDKVolumeResponse)
	return ret0
}

// InspectVolume indicates an expected call of InspectVolume
func (mr *MockDockerClientMockRecorder) InspectVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectVolume", reflect.TypeOf((*MockDockerClient)(nil).InspectVolume), arg0, arg1, arg2)
}

// KnownVersions mocks base method
func (m *MockDockerClient) KnownVersions() []dockerclient.DockerVersion {
	ret := m.ctrl.Call(m, "KnownVersions")
	ret0, _ := ret[0].([]dockerclient.DockerVersion)
	return ret0
}

// KnownVersions indicates an expected call of KnownVersions
func (mr *MockDockerClientMockRecorder) KnownVersions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KnownVersions", reflect.TypeOf((*MockDockerClient)(nil).KnownVersions))
}

// ListContainers mocks base method
func (m *MockDockerClient) ListContainers(arg0 context.Context, arg1 bool, arg2 time.Duration) dockerapi.ListContainersResponse {
	ret := m.ctrl.Call(m, "ListContainers", arg0, arg1, arg2)
	ret0, _ := ret[0].(dockerapi.ListContainersResponse)
	return ret0
}

// ListContainers indicates an expected call of ListContainers
func (mr *MockDockerClientMockRecorder) ListContainers(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockDockerClient)(nil).ListContainers), arg0, arg1, arg2)
}

// ListImages mocks base method
func (m *MockDockerClient) ListImages(arg0 context.Context, arg1 time.Duration) dockerapi.ListImagesResponse {
	ret := m.ctrl.Call(m, "ListImages", arg0, arg1)
	ret0, _ := ret[0].(dockerapi.ListImagesResponse)
	return ret0
}

// ListImages indicates an expected call of ListImages
func (mr *MockDockerClientMockRecorder) ListImages(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockDockerClient)(nil).ListImages), arg0, arg1)
}

// ListPlugins mocks base method
func (m *MockDockerClient) ListPlugins(arg0 context.Context, arg1 time.Duration, arg2 filters.Args) dockerapi.ListPluginsResponse {
	ret := m.ctrl.Call(m, "ListPlugins", arg0, arg1, arg2)
	ret0, _ := ret[0].(dockerapi.ListPluginsResponse)
	return ret0
}

// ListPlugins indicates an expected call of ListPlugins
func (mr *MockDockerClientMockRecorder) ListPlugins(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlugins", reflect.TypeOf((*MockDockerClient)(nil).ListPlugins), arg0, arg1, arg2)
}

// ListPluginsWithFilters mocks base method
func (m *MockDockerClient) ListPluginsWithFilters(arg0 context.Context, arg1 bool, arg2 []string, arg3 time.Duration) ([]string, error) {
	ret := m.ctrl.Call(m, "ListPluginsWithFilters", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPluginsWithFilters indicates an expected call of ListPluginsWithFilters
func (mr *MockDockerClientMockRecorder) ListPluginsWithFilters(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPluginsWithFilters", reflect.TypeOf((*MockDockerClient)(nil).ListPluginsWithFilters), arg0, arg1, arg2, arg3)
}

// LoadImage mocks base method
func (m *MockDockerClient) LoadImage(arg0 context.Context, arg1 io.Reader, arg2 time.Duration) error {
	ret := m.ctrl.Call(m, "LoadImage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadImage indicates an expected call of LoadImage
func (mr *MockDockerClientMockRecorder) LoadImage(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadImage", reflect.TypeOf((*MockDockerClient)(nil).LoadImage), arg0, arg1, arg2)
}

// PullImage mocks base method
func (m *MockDockerClient) PullImage(arg0 context.Context, arg1 string, arg2 *container.RegistryAuthenticationData, arg3 time.Duration) dockerapi.DockerContainerMetadata {
	ret := m.ctrl.Call(m, "PullImage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(dockerapi.DockerContainerMetadata)
	return ret0
}

// PullImage indicates an expected call of PullImage
func (mr *MockDockerClientMockRecorder) PullImage(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullImage", reflect.TypeOf((*MockDockerClient)(nil).PullImage), arg0, arg1, arg2, arg3)
}

// RemoveContainer mocks base method
func (m *MockDockerClient) RemoveContainer(arg0 context.Context, arg1 string, arg2 time.Duration) error {
	ret := m.ctrl.Call(m, "RemoveContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainer indicates an expected call of RemoveContainer
func (mr *MockDockerClientMockRecorder) RemoveContainer(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainer", reflect.TypeOf((*MockDockerClient)(nil).RemoveContainer), arg0, arg1, arg2)
}

// RemoveImage mocks base method
func (m *MockDockerClient) RemoveImage(arg0 context.Context, arg1 string, arg2 time.Duration) error {
	ret := m.ctrl.Call(m, "RemoveImage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveImage indicates an expected call of RemoveImage
func (mr *MockDockerClientMockRecorder) RemoveImage(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveImage", reflect.TypeOf((*MockDockerClient)(nil).RemoveImage), arg0, arg1, arg2)
}

// RemoveVolume mocks base method
func (m *MockDockerClient) RemoveVolume(arg0 context.Context, arg1 string, arg2 time.Duration) error {
	ret := m.ctrl.Call(m, "RemoveVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveVolume indicates an expected call of RemoveVolume
func (mr *MockDockerClientMockRecorder) RemoveVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVolume", reflect.TypeOf((*MockDockerClient)(nil).RemoveVolume), arg0, arg1, arg2)
}

// StartContainer mocks base method
func (m *MockDockerClient) StartContainer(arg0 context.Context, arg1 string, arg2 time.Duration) dockerapi.DockerContainerMetadata {
	ret := m.ctrl.Call(m, "StartContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].(dockerapi.DockerContainerMetadata)
	return ret0
}

// StartContainer indicates an expected call of StartContainer
func (mr *MockDockerClientMockRecorder) StartContainer(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContainer", reflect.TypeOf((*MockDockerClient)(nil).StartContainer), arg0, arg1, arg2)
}

// Stats mocks base method
func (m *MockDockerClient) Stats(arg0 context.Context, arg1 string, arg2 time.Duration) (<-chan *types.Stats, error) {
	ret := m.ctrl.Call(m, "Stats", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan *types.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (mr *MockDockerClientMockRecorder) Stats(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockDockerClient)(nil).Stats), arg0, arg1, arg2)
}

// StopContainer mocks base method
func (m *MockDockerClient) StopContainer(arg0 context.Context, arg1 string, arg2 time.Duration) dockerapi.DockerContainerMetadata {
	ret := m.ctrl.Call(m, "StopContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].(dockerapi.DockerContainerMetadata)
	return ret0
}

// StopContainer indicates an expected call of StopContainer
func (mr *MockDockerClientMockRecorder) StopContainer(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopContainer", reflect.TypeOf((*MockDockerClient)(nil).StopContainer), arg0, arg1, arg2)
}

// SupportedVersions mocks base method
func (m *MockDockerClient) SupportedVersions() []dockerclient.DockerVersion {
	ret := m.ctrl.Call(m, "SupportedVersions")
	ret0, _ := ret[0].([]dockerclient.DockerVersion)
	return ret0
}

// SupportedVersions indicates an expected call of SupportedVersions
func (mr *MockDockerClientMockRecorder) SupportedVersions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportedVersions", reflect.TypeOf((*MockDockerClient)(nil).SupportedVersions))
}

// Version mocks base method
func (m *MockDockerClient) Version(arg0 context.Context, arg1 time.Duration) (string, error) {
	ret := m.ctrl.Call(m, "Version", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockDockerClientMockRecorder) Version(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockDockerClient)(nil).Version), arg0, arg1)
}

// WithVersion mocks base method
func (m *MockDockerClient) WithVersion(arg0 dockerclient.DockerVersion) dockerapi.DockerClient {
	ret := m.ctrl.Call(m, "WithVersion", arg0)
	ret0, _ := ret[0].(dockerapi.DockerClient)
	return ret0
}

// WithVersion indicates an expected call of WithVersion
func (mr *MockDockerClientMockRecorder) WithVersion(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithVersion", reflect.TypeOf((*MockDockerClient)(nil).WithVersion), arg0)
}
