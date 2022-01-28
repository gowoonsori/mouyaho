package domain

import "context"

type BadgeId uint64

type BadgeService interface {
	renderLikeBadge(url string) ([]byte, error)
	GetBadge(ctx context.Context, url string) *Badge
	ParsingUrl(url string) (map[string]string, error)
}

type BadgeRepository interface {
	Save(b *Badge) (*Badge, error)
	FindById(id BadgeId) (*Badge, error)
	FindByUrl(url string) (*Badge, error)
}

type Badge struct {
	id       BadgeId
	url      string
	callUser UserId
	file     []byte
	reactCnt int
	isReact  bool
}

func NewBadge(url string, callUser UserId, file []byte, reactCnt int, isReact bool) *Badge {
	return &Badge{url: url, callUser: callUser, file: file, reactCnt: reactCnt, isReact: isReact}
}

func (b Badge) ReactBy() *React {
	return ByOn(b.callUser, b.id)
}

func (b Badge) UnReactBy() *React {
	return ByOn(b.callUser, b.id)
}
