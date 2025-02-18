// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	ara "github.com/trento-project/trento/web/services/ara"
)

// AraService is an autogenerated mock type for the AraService type
type AraService struct {
	mock.Mock
}

// GetRecord provides a mock function with given fields: recordId
func (_m *AraService) GetRecord(recordId int) (*ara.Record, error) {
	ret := _m.Called(recordId)

	var r0 *ara.Record
	if rf, ok := ret.Get(0).(func(int) *ara.Record); ok {
		r0 = rf(recordId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ara.Record)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(recordId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecordList provides a mock function with given fields: filter
func (_m *AraService) GetRecordList(filter string) (*ara.RecordList, error) {
	ret := _m.Called(filter)

	var r0 *ara.RecordList
	if rf, ok := ret.Get(0).(func(string) *ara.RecordList); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ara.RecordList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
