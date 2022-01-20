package user

type UserId uint64

type User struct {
	Id   UserId
	Name string
}

func NewUser(id UserId, name string) *User {
	return &User{Id: id, Name: name}
}
