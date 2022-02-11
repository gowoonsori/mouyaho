package application

import (
	"likeIt/domain"
	"likeIt/infrastructure/badge"
)

type LikeBadgeService struct {
	rr domain.ReactRepository
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
