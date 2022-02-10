// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	domain "likeIt/domain"

	mock "github.com/stretchr/testify/mock"
)

// BadgeService is an autogenerated mock type for the BadgeService type
type BadgeService struct {
	mock.Mock
}

// GetBadgeFile provides a mock function with given fields: id, url
func (_m *BadgeService) GetBadgeFile(id domain.UserId, url string) []byte {
	ret := _m.Called(id, url)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(domain.UserId, string) []byte); ok {
		r0 = rf(id, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}
