package badge

import "likeIt/domain/user"

type BadgeId uint64

type Badge struct {
	id      BadgeId
	url     string
	encoded string
	likers  []user.UserId
}

func NewBadge(id BadgeId, url string, encoded string, likers []user.UserId) *Badge {
	return &Badge{id: id, url: url, encoded: encoded, likers: likers}
}

func (b *Badge) IsLike(searchId user.UserId) bool {
	for _, likeId := range b.likers {
		if likeId == searchId {
			return true
		}
	}
	return false
}

func (b *Badge) GetLikeCount() int {
	return len(b.likers)
}

func (b *Badge) GetId() BadgeId {
	return b.id
}

func (b *Badge) GetLikers() []user.UserId {
	return b.likers
}

func (b *Badge) GetUrl() string {
	return b.url
}

func (b *Badge) GetEncoded() string {
	return b.encoded
}
