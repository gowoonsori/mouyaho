package domain

type UserId string

type User struct {
	Id   UserId
	Name string
}

type UserRepository interface {
	Save(b *User) (*User, error)
	FindById(id UserId) (*User, error)
}
