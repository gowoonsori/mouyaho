package domain

type UserId string

type User struct {
	Id UserId
}

type UserRepository interface {
	Save(b *User) (*User, error)
	FindById(id UserId) (*User, error)
}
