package dto

import (
	"likeIt/domain/badge"
	"likeIt/domain/user"
)

type LikeRequest struct {
	badge.BadgeId
	user.UserId
}