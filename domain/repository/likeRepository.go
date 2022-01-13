package repository

import (
	"likeIt/domain/aggregate"
	"likeIt/domain/aggregate/vo"
)

type LikeRepository interface {
	Like(like aggregate.Like) error
	Likes(likes []aggregate.Like) (int, error)
	UnLike(like aggregate.Like) error
	UnLikes(likes []aggregate.Like) (int, error)
	GetLikeCountByBadge(badge vo.Badge) (uint64, error)
	GetLikeByUser(user vo.User) (aggregate.Like, error)
}
