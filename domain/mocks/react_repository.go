// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	domain "likeIt/domain"

	mock "github.com/stretchr/testify/mock"
)

// ReactRepository is an autogenerated mock type for the ReactRepository type
type ReactRepository struct {
	mock.Mock
}

// DeleteById provides a mock function with given fields: id
func (_m *ReactRepository) DeleteById(id domain.ReactId) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.ReactId) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByBadgeId provides a mock function with given fields: badgeId
func (_m *ReactRepository) FindByBadgeId(badgeId domain.BadgeId) []domain.React {
	ret := _m.Called(badgeId)

	var r0 []domain.React
	if rf, ok := ret.Get(0).(func(domain.BadgeId) []domain.React); ok {
		r0 = rf(badgeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.React)
		}
	}

	return r0
}

// FindByBadgeIdAndUserId provides a mock function with given fields: badgeId, id
func (_m *ReactRepository) FindByBadgeIdAndUserId(badgeId domain.BadgeId, id domain.UserId) *domain.React {
	ret := _m.Called(badgeId, id)

	var r0 *domain.React
	if rf, ok := ret.Get(0).(func(domain.BadgeId, domain.UserId) *domain.React); ok {
		r0 = rf(badgeId, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.React)
		}
	}

	return r0
}

// FindCountByBadgeId provides a mock function with given fields: badgeId
func (_m *ReactRepository) FindCountByBadgeId(badgeId domain.BadgeId) int {
	ret := _m.Called(badgeId)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.BadgeId) int); ok {
		r0 = rf(badgeId)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Save provides a mock function with given fields: b
func (_m *ReactRepository) Save(b *domain.React) (*domain.React, error) {
	ret := _m.Called(b)

	var r0 *domain.React
	if rf, ok := ret.Get(0).(func(*domain.React) *domain.React); ok {
		r0 = rf(b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.React)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.React) error); ok {
		r1 = rf(b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLikeStatusById provides a mock function with given fields: id, status
func (_m *ReactRepository) UpdateLikeStatusById(id domain.ReactId, status bool) error {
	ret := _m.Called(id, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.ReactId, bool) error); ok {
		r0 = rf(id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
