package badge

import (
	"likeIt/domain/react"
	"likeIt/domain/user"
)

type BadgeId uint64

type Badge struct {
	id        BadgeId
	url       string
	badgeInfo BadgeInfo
}

func New(id BadgeId, url string) *Badge {
	return &Badge{id: id, url: url}
}

func (b *Badge) ReactBy(userId user.UserId) *react.React {
	return react.ByOn(userId, b.id)
}

func (b *Badge) UnReactBy(userId user.UserId) *react.React {
	return react.ByOn(userId, b.id)
}
