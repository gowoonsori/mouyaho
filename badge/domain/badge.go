package domain

type BadgeId string //url 이 BadgeId

type BadgeService interface {
	GetBadge(id UserId, url string) *Badge
	renderBadge() ([]byte, error)
}

type Badge struct {
	id        BadgeId
	badgeInfo BadgeInfo
}

func NewBadge(id BadgeId, badgeInfo BadgeInfo) *Badge {
	return &Badge{id: id, badgeInfo: badgeInfo}
}

func (b Badge) ReactBy(user UserId) *React {
	return ByOn(user, b.id)
}

func (b Badge) UnReactBy(user UserId) *React {
	return ByOn(user, b.id)
}

func (b Badge) Id() BadgeId {
	return b.id
}
