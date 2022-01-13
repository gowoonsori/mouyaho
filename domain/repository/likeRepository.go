package repository

import (
	"likeIt/domain/aggregate"
)

type LikeRepository interface {
	Like(like *aggregate.Like) error
	Likes(likes []aggregate.Like) (int, error)
	UnLike(like *aggregate.Like) error
	UnLikes(likes []aggregate.Like) (int, error)
	GetLikeCountByBadgeId(id uint64) (int, error)
	GetLikeByUserId(id uint64) (*aggregate.Like, error)
}
