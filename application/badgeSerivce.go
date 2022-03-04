package application

import (
	"likeIt/domain"
	"likeIt/infrastructure/badge"
)

type LikeBadgeService struct {
	rr domain.ReactRepository
}

func NewLikeBadgeService(rr domain.ReactRepository) *LikeBadgeService {
	return &LikeBadgeService{rr: rr}
}

func (lbs LikeBadgeService) GetBadgeFile(userId domain.UserId, reqUrl string) []byte {
	//query string parsing
	urlInfo := badge.CreateUrlInfoFromUrl(reqUrl)
	u := urlInfo.Url

	//like count get
	likeCount := lbs.rr.FindCountByBadgeId(domain.BadgeId(u))

	//isLike get
	il := lbs.rr.FindByBadgeIdAndUserId(domain.BadgeId(u), domain.UserId(userId))
	var isLike bool
	if il != nil {
		isLike = true
	}

	//file redering
	f, err := badge.GenerateLikeBadge(*urlInfo, isLike, likeCount)
	if err != nil {
		return []byte{}
	}

	return f
}

func (lbs LikeBadgeService) LikeBadge(userId domain.UserId, badgeId domain.BadgeId) *ReactDto {
	react := lbs.rr.FindByBadgeIdAndUserId(badgeId, userId)

	switch {
	case react != nil && react.IsLike():
		return nil
	case react != nil:
		if lbs.rr.UpdateLikeStatusById(react.Id(), true) != nil {
			return nil
		}
	case react == nil:
		if _, err := lbs.rr.Save(domain.ByOn(userId, badgeId)); err != nil {
			return nil
		}
	}

	cnt := lbs.rr.FindCountByBadgeId(badgeId)
	return &ReactDto{Count: cnt, IsLike: true}
}
func (lbs LikeBadgeService) ReactBadge(userId domain.UserId, badgeId domain.BadgeId) *ReactDto {
	react := lbs.rr.FindByBadgeIdAndUserId(badgeId, userId)

	if react != nil && react.IsLike() {
		if lbs.rr.UpdateLikeStatusById(react.Id(), false) != nil {
			return nil
		}
	} else {
		return nil
	}

	cnt := lbs.rr.FindCountByBadgeId(badgeId)
	return &ReactDto{Count: cnt, IsLike: true}
}
