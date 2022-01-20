package like

type Repository interface {
	Save(b *Like) (*Like, error)
	FindById(id LikeId) (*Like, error)
}
