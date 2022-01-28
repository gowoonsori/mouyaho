package domain

type UserId uint64

type User struct {
	Id   UserId
	Name string
}

type UserRepository interface {
	Save(b *User) (*User, error)
	FindById(id UserId) (*User, error)
}
