package react

import (
	"likeIt/domain/badge"
	"likeIt/domain/user"
)

type Repository interface {
	Save(b *React) (*React, error)
	FindByBadgeId(badgeId badge.BadgeId) []React
	FindByBadgeIdAndUserId(badgeId badge.BadgeId, id user.UserId) *React
	DeleteById(id ReactId) error
	FindCountByBadgeId(badgeId badge.BadgeId) int
}
