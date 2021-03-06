// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

import reportHealCheckDomain "github.com/oommi04/ReportHealcheck/domains/reportHealCheckDomain"

// ReportClientInterface is an autogenerated mock type for the ReportClientInterface type
type ReportClientInterface struct {
	mock.Mock
}

// ReportHealCheck provides a mock function with given fields: r
func (_m *ReportClientInterface) ReportHealCheck(r reportHealCheckDomain.ReportHealCheck) error {
	ret := _m.Called(r)

	var r0 error
	if rf, ok := ret.Get(0).(func(reportHealCheckDomain.ReportHealCheck) error); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
