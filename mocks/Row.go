// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Row is an autogenerated mock type for the Row type
type Row struct {
	mock.Mock
}

// Next provides a mock function with given fields:
func (_m *Row) Next() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Scan provides a mock function with given fields: dest
func (_m *Row) Scan(dest ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, dest...)
	_m.Called(_ca...)
}
