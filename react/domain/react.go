package domain

import (
	badge "likeIt/badge/domain"
	user "likeIt/user/domain"
)

type ReactId uint64

type React struct {
	id      ReactId
	badgeId badge.BadgeId
	userId  user.UserId
}

func (l *React) IsEmpty() bool {
	return l.id == 0 && l.badgeId == 0 && l.userId == 0
}

func ByOn(userId user.UserId, badgeId badge.BadgeId) *React {
	return &React{
		badgeId: badgeId,
		userId:  userId,
	}
}

type ReactRepository interface {
	Save(b *React) (*React, error)
	FindByBadgeId(badgeId badge.BadgeId) []React
	FindByBadgeIdAndUserId(badgeId badge.BadgeId, id user.UserId) *React
	DeleteById(id ReactId) error
	FindCountByBadgeId(badgeId badge.BadgeId) int
}
