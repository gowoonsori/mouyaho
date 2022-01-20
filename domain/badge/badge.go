package badge

import "likeIt/domain/like"

type BadgeId uint64
type ns struct{}

type Badge struct {
	Id      BadgeId
	Url     string
	Encoded string
	Likes   map[like.LikeId]ns
}

func New(id BadgeId, url string, encoded string) *Badge {
	return &Badge{Id: id, Url: url, Encoded: encoded}
}

func (b *Badge) IsLike(id like.LikeId) bool {
	return b.Likes[id] != ns{}
}

func (b *Badge) CountLike() int {
	return len(b.Likes)
}