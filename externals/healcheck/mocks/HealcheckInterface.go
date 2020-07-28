// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import time "time"
import url "net/url"

// HealcheckInterface is an autogenerated mock type for the HealcheckInterface type
type HealcheckInterface struct {
	mock.Mock
}

// HealCheckWebsite provides a mock function with given fields: _a0
func (_m *HealcheckInterface) HealCheckWebsite(_a0 *url.URL) (time.Duration, int, error) {
	ret := _m.Called(_a0)

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func(*url.URL) time.Duration); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*url.URL) int); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*url.URL) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
