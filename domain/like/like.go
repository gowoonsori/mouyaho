package like

import (
	"likeIt/domain/badge"
	"likeIt/domain/user"
)

type LikeId uint64

type Like struct {
	Id      LikeId
	BadgeId badge.BadgeId
	UserId  user.UserId
}

func New(id LikeId, badgeId badge.BadgeId, userId user.UserId) *Like {
	return &Like{Id: id, BadgeId: badgeId, UserId: userId}
}

func (l *Like) IsEmpty() bool {
	return l.Id == 0 && l.BadgeId == 0 && l.UserId == 0
}
