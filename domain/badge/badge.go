package badge

import "likeIt/domain/user"

type BadgeId uint64

type Badge struct {
	BadgeId
	likerList []user.UserId
}

func (b *Badge) IsLike(searchId user.UserId) bool {
	for _, likeId := range b.likerList {
		if likeId == searchId {
			return true
		}
	}
	return false
}

func (b *Badge) GetLikeCount() int {
	return len(b.likerList)
}
