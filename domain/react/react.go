package react

import (
	"likeIt/domain/badge"
	"likeIt/domain/user"
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
