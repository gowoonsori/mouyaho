package dto

import (
	"likeIt/domain/aggregate"
	"likeIt/domain/aggregate/vo"
)

type LikeRequest struct {
	BadgeId uint64
	UserId  uint64
}

func (lr *LikeRequest) CreateLike() aggregate.Like {
	return aggregate.Like{
		Badge:  vo.Badge{Id: lr.BadgeId},
		UserId: vo.User{Id: lr.UserId},
	}
}
