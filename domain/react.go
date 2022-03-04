package domain

type ReactId uint64

//go:generate mockery --name ReactRepository --case underscore
type ReactRepository interface {
	Save(b *React) (*React, error)
	FindByBadgeId(badgeId BadgeId) []React
	FindByBadgeIdAndUserId(badgeId BadgeId, id UserId) *React
	DeleteById(id ReactId) error
	FindCountByBadgeId(badgeId BadgeId) int
	UpdateLikeStatusById(id ReactId, status bool) error
}

type React struct {
	id      ReactId
	badgeId BadgeId
	reader  Reader
	isLike  bool
}

func NewReact(id ReactId, badgeId BadgeId, reader Reader, isLike bool) *React {
	return &React{id: id, badgeId: badgeId, reader: reader, isLike: isLike}
}

func ByOn(userId UserId, badgeId BadgeId) *React {
	return &React{
		badgeId: badgeId,
		reader:  *NewReader(userId),
	}
}

func (r React) Id() ReactId {
	return r.id
}

func (r React) IsLike() bool {
	return r.isLike
}
