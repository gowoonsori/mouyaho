package domain

type BadgeId string //url 이 BadgeId

//go:generate mockery --name BadgeService --case underscore
type BadgeService interface {
	GetBadgeFile(id UserId, url string) []byte
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
