// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import pager "github.com/kyma-project/kyma/components/ui-api-layer/internal/pager"
import resource "github.com/kyma-project/kyma/components/ui-api-layer/pkg/resource"

import v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"

// clusterServiceBrokerSvc is an autogenerated mock type for the clusterServiceBrokerSvc type
type clusterServiceBrokerSvc struct {
	mock.Mock
}

// Find provides a mock function with given fields: name
func (_m *clusterServiceBrokerSvc) Find(name string) (*v1beta1.ClusterServiceBroker, error) {
	ret := _m.Called(name)

	var r0 *v1beta1.ClusterServiceBroker
	if rf, ok := ret.Get(0).(func(string) *v1beta1.ClusterServiceBroker); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ClusterServiceBroker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: pagingParams
func (_m *clusterServiceBrokerSvc) List(pagingParams pager.PagingParams) ([]*v1beta1.ClusterServiceBroker, error) {
	ret := _m.Called(pagingParams)

	var r0 []*v1beta1.ClusterServiceBroker
	if rf, ok := ret.Get(0).(func(pager.PagingParams) []*v1beta1.ClusterServiceBroker); ok {
		r0 = rf(pagingParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.ClusterServiceBroker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pager.PagingParams) error); ok {
		r1 = rf(pagingParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: listener
func (_m *clusterServiceBrokerSvc) Subscribe(listener resource.Listener) {
	_m.Called(listener)
}

// Unsubscribe provides a mock function with given fields: listener
func (_m *clusterServiceBrokerSvc) Unsubscribe(listener resource.Listener) {
	_m.Called(listener)
}
