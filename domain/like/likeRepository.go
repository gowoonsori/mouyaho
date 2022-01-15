package like

type LikeRepository interface {
	Like(like *Like) error
	Likes(likes []Like) (int, error)
	UnLike(like *Like) error
	UnLikes(likes []Like) (int, error)
	GetLikeCountByBadgeId(id uint64) (int, error)
	GetLikeByUserId(id uint64) (*Like, error)
}
