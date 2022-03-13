package infrastructure

import (
	domain "likeIt/domain"

	mock "github.com/stretchr/testify/mock"
)

type BadgeRepository struct {
	mock.Mock
}

func (_m *BadgeRepository) GetReactionsByBadge(_a0 *domain.Badge) *domain.Badge {
	ret := _m.Called(_a0)

	var r0 *domain.Badge
	if rf, ok := ret.Get(0).(func(*domain.Badge) *domain.Badge); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Badge)
		}
	}

	return r0
}
