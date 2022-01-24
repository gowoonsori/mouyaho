package domain

import (
	"context"
)

type BadgeId uint64

type Badge struct {
	id       BadgeId
	url      string
	callUser UserId
	icon     byte[]
	reactCnt int
	isReact  bool
}

func (b Badge) ReactBy() *React {
	return ByOn(b.callUser, b.id)
}

func (b Badge) UnReactBy() *React {
	return ByOn(b.callUser, b.id)
}

type BadgeUsecase interface {
	GetBadge(ctx context.Context, url string) *Badge
}

type BadgeRepository interface {
	Save(b *Badge) (*Badge, error)
	FindById(id BadgeId) (*Badge, error)
	FindByUrl(url string) (*Badge, error)
}
