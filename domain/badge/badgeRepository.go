package badge

type Repository interface {
	Save(b *Badge) (*Badge, error)
	FindById(id BadgeId) (*Badge, error)
}
