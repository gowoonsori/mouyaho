package application

import (
	domain "likeIt/domain"

	mock "github.com/stretchr/testify/mock"
)

type BadgeService struct {
	mock.Mock
}

func (_m *BadgeService) RenderBadgeHtml(badge domain.Badge) []byte {
	ret := _m.Called(badge)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(domain.Badge) []byte); ok {
		r0 = rf(badge)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}
