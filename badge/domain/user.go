package domain

type UserId string

type user struct {
	id UserId
}

func (u *user) Id() UserId {
	return u.id
}

type Reader struct {
	user
}

func NewReader(userId UserId) *Reader {
	return &Reader{user: user{
		id: userId,
	}}
}
