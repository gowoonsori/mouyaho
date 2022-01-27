package domain

import (
	"context"
	react "likeIt/react/domain"
	user "likeIt/user/domain"
)

type BadgeId uint64

type Badge struct {
	id       BadgeId
	url      string
	callUser user.UserId
	file     byte[]
	reactCnt int
	isReact  bool
}

func (b Badge) ReactBy() *react.React {
	return react.ByOn(b.callUser, b.id)
}

func (b Badge) UnReactBy() *react.React {
	return react.ByOn(b.callUser, b.id)
}

type BadgeUsecase interface {
	renderBadge(ctx context.Context) ([]byte, error)
	GetBadge(ctx context.Context, url string) *Badge
}

type BadgeRepository interface {
	Save(b *Badge) (*Badge, error)
	FindById(id BadgeId) (*Badge, error)
	FindByUrl(url string) (*Badge, error)
}
