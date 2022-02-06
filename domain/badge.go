package domain

type BadgeId string //url 이 BadgeId

type BadgeService interface {
	GetBadge(id UserId, url string) *Badge
	renderBadge() ([]byte, error)
}

//go:generate mockery --name BadgeRepository --case underscore
type BadgeRepository interface {
	Save(b *Badge) (*Badge, error)
	FindById(id BadgeId) (*Badge, error)
}

type Badge struct {
	id   BadgeId
	file []byte
}

func NewBadge(id BadgeId, file []byte) *Badge {
	return &Badge{id: id, file: file}
}

func (b Badge) ReactBy(user UserId) *React {
	return ByOn(user, b.id)
}

func (b Badge) UnReactBy(user UserId) *React {
	return ByOn(user, b.id)
}

func (b Badge) Id() BadgeId {
	return b.id
}

func (b Badge) File() []byte {
	return b.file
}
