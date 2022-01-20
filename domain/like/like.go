package like

import (
	"likeIt/domain/badge"
	"likeIt/domain/user"
)

type ns struct{}
type LikeId uint64

type Like struct {
	Id      LikeId
	BadgeId badge.BadgeId
	Likers  map[user.UserId]ns
}

func New(id LikeId, badgeId badge.BadgeId, likers map[user.UserId]ns) *Like {
	return &Like{Id: id, BadgeId: badgeId, Likers: likers}
}

func (l *Like) AddLiker(id user.UserId) {
	l.Likers[id] = ns{}
}

func (l *Like) DeleteLiker(id user.UserId) {
	delete(l.Likers, id)
}

func (l *Like) GetLikerCount() int {
	return len(l.Likers)
}

func (l *Like) IsEmpty() bool {
	return l.Id == 0 && l.BadgeId == 0 && len(l.Likers) == 0
}
