package like

import (
	"likeIt/domain/badge"
	"likeIt/domain/user"
)

type Repository interface {
	Save(b *Like) (*Like, error)
	FindByBadgeId(badgeId badge.BadgeId) []Like
	FindByBadgeIdAndUserId(badgeId badge.BadgeId, id user.UserId) *Like
	DeleteById(id LikeId) error
}
