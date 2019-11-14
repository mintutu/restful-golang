// Code generated by mockery v1.0.0. DO NOT EDIT.

package service

import mock "github.com/stretchr/testify/mock"

// MockHTTPClient is an autogenerated mock type for the HTTPClient type
type MockHTTPClient struct {
	mock.Mock
}

// GetData provides a mock function with given fields: _a0
func (_m *MockHTTPClient) GetData(_a0 string) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
