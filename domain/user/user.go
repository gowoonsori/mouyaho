package user

import "likeIt/domain/badge"

type UserId uint64

type User struct {
	id UserId
}

func New(id UserId, likes []badge.BadgeId) *User {
	return &User{id: id}
}

func (u User) GetId() UserId {
	return u.id
}
