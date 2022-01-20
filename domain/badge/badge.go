package badge

type BadgeId uint64

type Badge struct {
	Id        BadgeId
	Url       string
	Encoded   string
	LikeCount int
}

func New(id BadgeId, url string, encoded string, likeCount int) *Badge {
	return &Badge{Id: id, Url: url, Encoded: encoded, LikeCount: likeCount}
}
