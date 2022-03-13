package domain

type UserId string

type user struct {
	id UserId
}

func (u *user) Id() UserId {
	return u.id
}

type Reactor struct {
	user
}

func NewReactor(userId UserId) *Reactor {
	return &Reactor{user: user{
		id: userId,
	}}
}
