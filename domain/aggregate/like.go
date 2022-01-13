package aggregate

import "likeIt/domain/aggregate/vo"

type Like struct {
	Id     uint64
	Badge  vo.Badge
	UserId vo.User
}
