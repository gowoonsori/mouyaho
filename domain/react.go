package domain

type ReactId uint64

type React struct {
	id      ReactId
	badgeId BadgeId
	userId  UserId
}

func (l *React) IsEmpty() bool {
	return l.id == 0 && l.badgeId == 0 && l.userId == 0
}

func ByOn(userId UserId, badgeId BadgeId) *React {
	return &React{
		badgeId: badgeId,
		userId:  userId,
	}
}

type ReactRepository interface {
	Save(b *React) (*React, error)
	FindByBadgeId(badgeId BadgeId) []React
	FindByBadgeIdAndUserId(badgeId BadgeId, id UserId) *React
	DeleteById(id ReactId) error
	FindCountByBadgeId(badgeId BadgeId) int
}
