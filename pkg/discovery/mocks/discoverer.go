// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "nimona.io/internal/context"
import discovery "nimona.io/pkg/discovery"
import mock "github.com/stretchr/testify/mock"
import peer "nimona.io/pkg/net/peer"

// Discoverer is an autogenerated mock type for the Discoverer type
type Discoverer struct {
	mock.Mock
}

// Add provides a mock function with given fields: v
func (_m *Discoverer) Add(v *peer.PeerInfo) {
	_m.Called(v)
}

// AddProvider provides a mock function with given fields: provider
func (_m *Discoverer) AddProvider(provider discovery.Provider) error {
	ret := _m.Called(provider)

	var r0 error
	if rf, ok := ret.Get(0).(func(discovery.Provider) error); ok {
		r0 = rf(provider)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Discover provides a mock function with given fields: ctx, q, options
func (_m *Discoverer) Discover(ctx context.Context, q *peer.PeerInfoRequest, options ...discovery.DiscovererOption) ([]*peer.PeerInfo, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, q)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*peer.PeerInfo
	if rf, ok := ret.Get(0).(func(context.Context, *peer.PeerInfoRequest, ...discovery.DiscovererOption) []*peer.PeerInfo); ok {
		r0 = rf(ctx, q, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*peer.PeerInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *peer.PeerInfoRequest, ...discovery.DiscovererOption) error); ok {
		r1 = rf(ctx, q, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
