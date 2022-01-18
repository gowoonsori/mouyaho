package badge

import "likeIt/domain/user"

type Repository interface {
	Add(b *Badge)
	FindById(id BadgeId) *Badge
	AddLiker(id user.UserId) error
	DeleteLiker(id user.UserId) error
	FindLikerList(id BadgeId) []user.UserId
}
