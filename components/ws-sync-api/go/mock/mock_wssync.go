// Copyright (c) 2020 TypeFox GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by MockGen. DO NOT EDIT.
// Source: wssync.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	api "github.com/gitpod-io/gitpod/ws-sync/api"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockWorkspaceContentServiceClient is a mock of WorkspaceContentServiceClient interface
type MockWorkspaceContentServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceContentServiceClientMockRecorder
}

// MockWorkspaceContentServiceClientMockRecorder is the mock recorder for MockWorkspaceContentServiceClient
type MockWorkspaceContentServiceClientMockRecorder struct {
	mock *MockWorkspaceContentServiceClient
}

// NewMockWorkspaceContentServiceClient creates a new mock instance
func NewMockWorkspaceContentServiceClient(ctrl *gomock.Controller) *MockWorkspaceContentServiceClient {
	mock := &MockWorkspaceContentServiceClient{ctrl: ctrl}
	mock.recorder = &MockWorkspaceContentServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorkspaceContentServiceClient) EXPECT() *MockWorkspaceContentServiceClientMockRecorder {
	return m.recorder
}

// InitWorkspace mocks base method
func (m *MockWorkspaceContentServiceClient) InitWorkspace(ctx context.Context, in *api.InitWorkspaceRequest, opts ...grpc.CallOption) (*api.InitWorkspaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InitWorkspace", varargs...)
	ret0, _ := ret[0].(*api.InitWorkspaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitWorkspace indicates an expected call of InitWorkspace
func (mr *MockWorkspaceContentServiceClientMockRecorder) InitWorkspace(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitWorkspace", reflect.TypeOf((*MockWorkspaceContentServiceClient)(nil).InitWorkspace), varargs...)
}

// WaitForInit mocks base method
func (m *MockWorkspaceContentServiceClient) WaitForInit(ctx context.Context, in *api.WaitForInitRequest, opts ...grpc.CallOption) (*api.WaitForInitResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitForInit", varargs...)
	ret0, _ := ret[0].(*api.WaitForInitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForInit indicates an expected call of WaitForInit
func (mr *MockWorkspaceContentServiceClientMockRecorder) WaitForInit(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForInit", reflect.TypeOf((*MockWorkspaceContentServiceClient)(nil).WaitForInit), varargs...)
}

// TakeSnapshot mocks base method
func (m *MockWorkspaceContentServiceClient) TakeSnapshot(ctx context.Context, in *api.TakeSnapshotRequest, opts ...grpc.CallOption) (*api.TakeSnapshotResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TakeSnapshot", varargs...)
	ret0, _ := ret[0].(*api.TakeSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TakeSnapshot indicates an expected call of TakeSnapshot
func (mr *MockWorkspaceContentServiceClientMockRecorder) TakeSnapshot(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeSnapshot", reflect.TypeOf((*MockWorkspaceContentServiceClient)(nil).TakeSnapshot), varargs...)
}

// DisposeWorkspace mocks base method
func (m *MockWorkspaceContentServiceClient) DisposeWorkspace(ctx context.Context, in *api.DisposeWorkspaceRequest, opts ...grpc.CallOption) (*api.DisposeWorkspaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisposeWorkspace", varargs...)
	ret0, _ := ret[0].(*api.DisposeWorkspaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisposeWorkspace indicates an expected call of DisposeWorkspace
func (mr *MockWorkspaceContentServiceClientMockRecorder) DisposeWorkspace(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisposeWorkspace", reflect.TypeOf((*MockWorkspaceContentServiceClient)(nil).DisposeWorkspace), varargs...)
}

// MockWorkspaceContentServiceServer is a mock of WorkspaceContentServiceServer interface
type MockWorkspaceContentServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceContentServiceServerMockRecorder
}

// MockWorkspaceContentServiceServerMockRecorder is the mock recorder for MockWorkspaceContentServiceServer
type MockWorkspaceContentServiceServerMockRecorder struct {
	mock *MockWorkspaceContentServiceServer
}

// NewMockWorkspaceContentServiceServer creates a new mock instance
func NewMockWorkspaceContentServiceServer(ctrl *gomock.Controller) *MockWorkspaceContentServiceServer {
	mock := &MockWorkspaceContentServiceServer{ctrl: ctrl}
	mock.recorder = &MockWorkspaceContentServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorkspaceContentServiceServer) EXPECT() *MockWorkspaceContentServiceServerMockRecorder {
	return m.recorder
}

// InitWorkspace mocks base method
func (m *MockWorkspaceContentServiceServer) InitWorkspace(arg0 context.Context, arg1 *api.InitWorkspaceRequest) (*api.InitWorkspaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitWorkspace", arg0, arg1)
	ret0, _ := ret[0].(*api.InitWorkspaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitWorkspace indicates an expected call of InitWorkspace
func (mr *MockWorkspaceContentServiceServerMockRecorder) InitWorkspace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitWorkspace", reflect.TypeOf((*MockWorkspaceContentServiceServer)(nil).InitWorkspace), arg0, arg1)
}

// WaitForInit mocks base method
func (m *MockWorkspaceContentServiceServer) WaitForInit(arg0 context.Context, arg1 *api.WaitForInitRequest) (*api.WaitForInitResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForInit", arg0, arg1)
	ret0, _ := ret[0].(*api.WaitForInitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForInit indicates an expected call of WaitForInit
func (mr *MockWorkspaceContentServiceServerMockRecorder) WaitForInit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForInit", reflect.TypeOf((*MockWorkspaceContentServiceServer)(nil).WaitForInit), arg0, arg1)
}

// TakeSnapshot mocks base method
func (m *MockWorkspaceContentServiceServer) TakeSnapshot(arg0 context.Context, arg1 *api.TakeSnapshotRequest) (*api.TakeSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakeSnapshot", arg0, arg1)
	ret0, _ := ret[0].(*api.TakeSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TakeSnapshot indicates an expected call of TakeSnapshot
func (mr *MockWorkspaceContentServiceServerMockRecorder) TakeSnapshot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeSnapshot", reflect.TypeOf((*MockWorkspaceContentServiceServer)(nil).TakeSnapshot), arg0, arg1)
}

// DisposeWorkspace mocks base method
func (m *MockWorkspaceContentServiceServer) DisposeWorkspace(arg0 context.Context, arg1 *api.DisposeWorkspaceRequest) (*api.DisposeWorkspaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisposeWorkspace", arg0, arg1)
	ret0, _ := ret[0].(*api.DisposeWorkspaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisposeWorkspace indicates an expected call of DisposeWorkspace
func (mr *MockWorkspaceContentServiceServerMockRecorder) DisposeWorkspace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisposeWorkspace", reflect.TypeOf((*MockWorkspaceContentServiceServer)(nil).DisposeWorkspace), arg0, arg1)
}
