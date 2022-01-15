package dto

import (
	"likeIt/domain/badge"
	"likeIt/domain/like"
	"likeIt/domain/user"
)

type LikeRequest struct {
	BadgeId uint64
	UserId  uint64
}

func (lr *LikeRequest) CreateLike() like.Like {
	return like.Like{
		Badge:  badge.Badge{Id: lr.BadgeId},
		UserId: user.User{Id: lr.UserId},
	}
}
