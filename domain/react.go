package domain

type ReactId uint64

//go:generate mockery --name ReactRepository --case underscore
type ReactRepository interface {
	Save(b *React) (*React, error)
	FindByBadgeId(badgeId BadgeId) []React
	FindByBadgeIdAndUserId(badgeId BadgeId, id UserId) *React
	DeleteById(id ReactId) error
	FindCountByBadgeId(badgeId BadgeId) int
}

type React struct {
	id      ReactId
	badgeId BadgeId
	userId  UserId
}

func NewReact(id ReactId, badgeId BadgeId, userId UserId) *React {
	return &React{id: id, badgeId: badgeId, userId: userId}
}

func ByOn(userId UserId, badgeId BadgeId) *React {
	return &React{
		badgeId: badgeId,
		userId:  userId,
	}
}
