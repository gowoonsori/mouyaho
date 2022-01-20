package user

type Repository interface {
	Save(b *User) (*User, error)
	FindById(id UserId) (*User, error)
}
