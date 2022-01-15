package like

import (
	"likeIt/domain/badge"
	"likeIt/domain/user"
)

type Like struct {
	Id     uint64
	Badge  badge.Badge
	UserId user.User
}
