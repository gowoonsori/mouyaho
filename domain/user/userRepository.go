package user

type Repository interface {
	Add(u *User) error
	GetById(id uint64) *User
}
