// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/kf/pkg/kf/services/fake (interfaces: ClientInterface)

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	services "github.com/google/kf/pkg/kf/services"
	v1beta1 "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	reflect "reflect"
)

// FakeClientInterface is a mock of ClientInterface interface
type FakeClientInterface struct {
	ctrl     *gomock.Controller
	recorder *FakeClientInterfaceMockRecorder
}

// FakeClientInterfaceMockRecorder is the mock recorder for FakeClientInterface
type FakeClientInterfaceMockRecorder struct {
	mock *FakeClientInterface
}

// NewFakeClientInterface creates a new mock instance
func NewFakeClientInterface(ctrl *gomock.Controller) *FakeClientInterface {
	mock := &FakeClientInterface{ctrl: ctrl}
	mock.recorder = &FakeClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *FakeClientInterface) EXPECT() *FakeClientInterfaceMockRecorder {
	return m.recorder
}

// CreateService mocks base method
func (m *FakeClientInterface) CreateService(arg0, arg1, arg2 string, arg3 ...services.CreateServiceOption) (*v1beta1.ServiceInstance, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateService", varargs...)
	ret0, _ := ret[0].(*v1beta1.ServiceInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService
func (mr *FakeClientInterfaceMockRecorder) CreateService(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*FakeClientInterface)(nil).CreateService), varargs...)
}

// DeleteService mocks base method
func (m *FakeClientInterface) DeleteService(arg0 string, arg1 ...services.DeleteServiceOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteService", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService
func (mr *FakeClientInterfaceMockRecorder) DeleteService(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*FakeClientInterface)(nil).DeleteService), varargs...)
}

// GetService mocks base method
func (m *FakeClientInterface) GetService(arg0 string, arg1 ...services.GetServiceOption) (*v1beta1.ServiceInstance, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetService", varargs...)
	ret0, _ := ret[0].(*v1beta1.ServiceInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *FakeClientInterfaceMockRecorder) GetService(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*FakeClientInterface)(nil).GetService), varargs...)
}

// ListServices mocks base method
func (m *FakeClientInterface) ListServices(arg0 ...services.ListServicesOption) (*v1beta1.ServiceInstanceList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*v1beta1.ServiceInstanceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *FakeClientInterfaceMockRecorder) ListServices(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*FakeClientInterface)(nil).ListServices), arg0...)
}

// Marketplace mocks base method
func (m *FakeClientInterface) Marketplace(arg0 ...services.MarketplaceOption) (*services.KfMarketplace, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Marketplace", varargs...)
	ret0, _ := ret[0].(*services.KfMarketplace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marketplace indicates an expected call of Marketplace
func (mr *FakeClientInterfaceMockRecorder) Marketplace(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marketplace", reflect.TypeOf((*FakeClientInterface)(nil).Marketplace), arg0...)
}
