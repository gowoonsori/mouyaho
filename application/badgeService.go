package application

import (
	"likeIt/domain"
	badgeInfra "likeIt/infrastructure/badge"
	"strconv"
)

type likeBadgeService struct {
	badgeRepository domain.BadgeRepository
}

func NewLikeBadgeService(badgeRepository domain.BadgeRepository) *likeBadgeService {
	return &likeBadgeService{badgeRepository: badgeRepository}
}

func (lbs likeBadgeService) CreateHeartBadge(ib domain.Badge, userId uint64) *badgeInfra.HeartBadge {
	badgeId := domain.BadgeId{
		Owner: ib.Owner,
		Repo:  ib.Repo,
		Title: ib.Title,
	}
	reactions := lbs.badgeRepository.GetReactionsBy(badgeId)
	b := domain.CreateBadge(ib.Owner, ib.Repo, ib.Title, ib.BgColor, ib.BorderColor, ib.IconColor, ib.ReactColor, ib.TextColor,
		ib.ShareColor, ib.Edge, reactions)

	var isReact bool
	if r := b.GetHeartReactionByUserId(userId); r != nil {
		isReact = true
	}

	cnt := b.GetHeartCount()

	hb := badgeInfra.NewHeartBadge(b.BgColor, b.BorderColor, b.IconColor, b.ReactColor, isReact, b.TextColor,
		strconv.Itoa(cnt), b.ShareColor, b.Edge)
	return hb
}
