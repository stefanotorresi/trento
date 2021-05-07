// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	api "github.com/hashicorp/consul/api"

	mock "github.com/stretchr/testify/mock"
)

// KV is an autogenerated mock type for the KV type
type KV struct {
	mock.Mock
}

// DeleteTree provides a mock function with given fields: prefix, w
func (_m *KV) DeleteTree(prefix string, w *api.WriteOptions) (*api.WriteMeta, error) {
	ret := _m.Called(prefix, w)

	var r0 *api.WriteMeta
	if rf, ok := ret.Get(0).(func(string, *api.WriteOptions) *api.WriteMeta); ok {
		r0 = rf(prefix, w)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.WriteMeta)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *api.WriteOptions) error); ok {
		r1 = rf(prefix, w)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: key, q
func (_m *KV) Get(key string, q *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error) {
	ret := _m.Called(key, q)

	var r0 *api.KVPair
	if rf, ok := ret.Get(0).(func(string, *api.QueryOptions) *api.KVPair); ok {
		r0 = rf(key, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.KVPair)
		}
	}

	var r1 *api.QueryMeta
	if rf, ok := ret.Get(1).(func(string, *api.QueryOptions) *api.QueryMeta); ok {
		r1 = rf(key, q)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*api.QueryMeta)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, *api.QueryOptions) error); ok {
		r2 = rf(key, q)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Keys provides a mock function with given fields: prefix, separator, q
func (_m *KV) Keys(prefix string, separator string, q *api.QueryOptions) ([]string, *api.QueryMeta, error) {
	ret := _m.Called(prefix, separator, q)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string, *api.QueryOptions) []string); ok {
		r0 = rf(prefix, separator, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 *api.QueryMeta
	if rf, ok := ret.Get(1).(func(string, string, *api.QueryOptions) *api.QueryMeta); ok {
		r1 = rf(prefix, separator, q)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*api.QueryMeta)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, *api.QueryOptions) error); ok {
		r2 = rf(prefix, separator, q)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// List provides a mock function with given fields: prefix, q
func (_m *KV) List(prefix string, q *api.QueryOptions) (api.KVPairs, *api.QueryMeta, error) {
	ret := _m.Called(prefix, q)

	var r0 api.KVPairs
	if rf, ok := ret.Get(0).(func(string, *api.QueryOptions) api.KVPairs); ok {
		r0 = rf(prefix, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.KVPairs)
		}
	}

	var r1 *api.QueryMeta
	if rf, ok := ret.Get(1).(func(string, *api.QueryOptions) *api.QueryMeta); ok {
		r1 = rf(prefix, q)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*api.QueryMeta)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, *api.QueryOptions) error); ok {
		r2 = rf(prefix, q)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListMap provides a mock function with given fields: prefix, offset
func (_m *KV) ListMap(prefix string, offset string) (map[string]interface{}, error) {
	ret := _m.Called(prefix, offset)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string, string) map[string]interface{}); ok {
		r0 = rf(prefix, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(prefix, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: p, q
func (_m *KV) Put(p *api.KVPair, q *api.WriteOptions) (*api.WriteMeta, error) {
	ret := _m.Called(p, q)

	var r0 *api.WriteMeta
	if rf, ok := ret.Get(0).(func(*api.KVPair, *api.WriteOptions) *api.WriteMeta); ok {
		r0 = rf(p, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.WriteMeta)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*api.KVPair, *api.WriteOptions) error); ok {
		r1 = rf(p, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutMap provides a mock function with given fields: prefix, data
func (_m *KV) PutMap(prefix string, data map[string]interface{}) error {
	ret := _m.Called(prefix, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) error); ok {
		r0 = rf(prefix, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutTyped provides a mock function with given fields: prefix, value
func (_m *KV) PutTyped(prefix string, value interface{}) error {
	ret := _m.Called(prefix, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(prefix, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
