// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
// Source: github.com/aws/amazon-vpc-cni-k8s/pkg/awsutils (interfaces: APIs)

// Package mock_awsutils is a generated GoMock package.
package mock_awsutils

import (
	reflect "reflect"

	awsutils "github.com/aws/amazon-vpc-cni-k8s/pkg/awsutils"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	gomock "github.com/golang/mock/gomock"
)

// MockAPIs is a mock of APIs interface
type MockAPIs struct {
	ctrl     *gomock.Controller
	recorder *MockAPIsMockRecorder
}

// MockAPIsMockRecorder is the mock recorder for MockAPIs
type MockAPIsMockRecorder struct {
	mock *MockAPIs
}

// NewMockAPIs creates a new mock instance
func NewMockAPIs(ctrl *gomock.Controller) *MockAPIs {
	mock := &MockAPIs{ctrl: ctrl}
	mock.recorder = &MockAPIsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPIs) EXPECT() *MockAPIsMockRecorder {
	return m.recorder
}

// AllocAllIPAddress mocks base method
func (m *MockAPIs) AllocAllIPAddress(arg0 string) error {
	ret := m.ctrl.Call(m, "AllocAllIPAddress", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AllocAllIPAddress indicates an expected call of AllocAllIPAddress
func (mr *MockAPIsMockRecorder) AllocAllIPAddress(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocAllIPAddress", reflect.TypeOf((*MockAPIs)(nil).AllocAllIPAddress), arg0)
}

// AllocENI mocks base method
func (m *MockAPIs) AllocENI(arg0 bool, arg1 []*string, arg2 string) (string, error) {
	ret := m.ctrl.Call(m, "AllocENI", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllocENI indicates an expected call of AllocENI
func (mr *MockAPIsMockRecorder) AllocENI(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocENI", reflect.TypeOf((*MockAPIs)(nil).AllocENI), arg0, arg1, arg2)
}

// AllocIPAddress mocks base method
func (m *MockAPIs) AllocIPAddress(arg0 string) error {
	ret := m.ctrl.Call(m, "AllocIPAddress", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AllocIPAddress indicates an expected call of AllocIPAddress
func (mr *MockAPIsMockRecorder) AllocIPAddress(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocIPAddress", reflect.TypeOf((*MockAPIs)(nil).AllocIPAddress), arg0)
}

// AllocIPAddresses mocks base method
func (m *MockAPIs) AllocIPAddresses(arg0 string, arg1 int64) error {
	ret := m.ctrl.Call(m, "AllocIPAddresses", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AllocIPAddresses indicates an expected call of AllocIPAddresses
func (mr *MockAPIsMockRecorder) AllocIPAddresses(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocIPAddresses", reflect.TypeOf((*MockAPIs)(nil).AllocIPAddresses), arg0, arg1)
}

// DescribeENI mocks base method
func (m *MockAPIs) DescribeENI(arg0 string) ([]*ec2.NetworkInterfacePrivateIpAddress, *string, error) {
	ret := m.ctrl.Call(m, "DescribeENI", arg0)
	ret0, _ := ret[0].([]*ec2.NetworkInterfacePrivateIpAddress)
	ret1, _ := ret[1].(*string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DescribeENI indicates an expected call of DescribeENI
func (mr *MockAPIsMockRecorder) DescribeENI(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeENI", reflect.TypeOf((*MockAPIs)(nil).DescribeENI), arg0)
}

// FreeENI mocks base method
func (m *MockAPIs) FreeENI(arg0 string) error {
	ret := m.ctrl.Call(m, "FreeENI", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FreeENI indicates an expected call of FreeENI
func (mr *MockAPIsMockRecorder) FreeENI(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreeENI", reflect.TypeOf((*MockAPIs)(nil).FreeENI), arg0)
}

// GetAttachedENIs mocks base method
func (m *MockAPIs) GetAttachedENIs() ([]awsutils.ENIMetadata, error) {
	ret := m.ctrl.Call(m, "GetAttachedENIs")
	ret0, _ := ret[0].([]awsutils.ENIMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttachedENIs indicates an expected call of GetAttachedENIs
func (mr *MockAPIsMockRecorder) GetAttachedENIs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachedENIs", reflect.TypeOf((*MockAPIs)(nil).GetAttachedENIs))
}

// GetENILimit mocks base method
func (m *MockAPIs) GetENILimit() (int, error) {
	ret := m.ctrl.Call(m, "GetENILimit")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetENILimit indicates an expected call of GetENILimit
func (mr *MockAPIsMockRecorder) GetENILimit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetENILimit", reflect.TypeOf((*MockAPIs)(nil).GetENILimit))
}

// GetENIipLimit mocks base method
func (m *MockAPIs) GetENIipLimit() (int64, error) {
	ret := m.ctrl.Call(m, "GetENIipLimit")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetENIipLimit indicates an expected call of GetENIipLimit
func (mr *MockAPIsMockRecorder) GetENIipLimit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetENIipLimit", reflect.TypeOf((*MockAPIs)(nil).GetENIipLimit))
}

// GetLocalIPv4 mocks base method
func (m *MockAPIs) GetLocalIPv4() string {
	ret := m.ctrl.Call(m, "GetLocalIPv4")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLocalIPv4 indicates an expected call of GetLocalIPv4
func (mr *MockAPIsMockRecorder) GetLocalIPv4() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocalIPv4", reflect.TypeOf((*MockAPIs)(nil).GetLocalIPv4))
}

// GetPrimaryENI mocks base method
func (m *MockAPIs) GetPrimaryENI() string {
	ret := m.ctrl.Call(m, "GetPrimaryENI")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPrimaryENI indicates an expected call of GetPrimaryENI
func (mr *MockAPIsMockRecorder) GetPrimaryENI() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrimaryENI", reflect.TypeOf((*MockAPIs)(nil).GetPrimaryENI))
}

// GetPrimaryENImac mocks base method
func (m *MockAPIs) GetPrimaryENImac() string {
	ret := m.ctrl.Call(m, "GetPrimaryENImac")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPrimaryENImac indicates an expected call of GetPrimaryENImac
func (mr *MockAPIsMockRecorder) GetPrimaryENImac() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrimaryENImac", reflect.TypeOf((*MockAPIs)(nil).GetPrimaryENImac))
}

// GetVPCIPv4CIDR mocks base method
func (m *MockAPIs) GetVPCIPv4CIDR() string {
	ret := m.ctrl.Call(m, "GetVPCIPv4CIDR")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVPCIPv4CIDR indicates an expected call of GetVPCIPv4CIDR
func (mr *MockAPIsMockRecorder) GetVPCIPv4CIDR() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVPCIPv4CIDR", reflect.TypeOf((*MockAPIs)(nil).GetVPCIPv4CIDR))
}

// GetVPCIPv4CIDRs mocks base method
func (m *MockAPIs) GetVPCIPv4CIDRs() []*string {
	ret := m.ctrl.Call(m, "GetVPCIPv4CIDRs")
	ret0, _ := ret[0].([]*string)
	return ret0
}

// GetVPCIPv4CIDRs indicates an expected call of GetVPCIPv4CIDRs
func (mr *MockAPIsMockRecorder) GetVPCIPv4CIDRs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVPCIPv4CIDRs", reflect.TypeOf((*MockAPIs)(nil).GetVPCIPv4CIDRs))
}
