package user

import "likeIt/domain/badge"

type UserId uint64

type User struct {
	Id UserId
}

func New(id UserId, likes []badge.BadgeId) *User {
	return &User{Id: id}
}
