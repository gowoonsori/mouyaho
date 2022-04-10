package domain

type Reaction struct {
	Id       int64
	UserId   int64
	UserName string
	Content  string
}

func (r Reaction) isHeart() bool {
	return r.Content == heartContent
}
