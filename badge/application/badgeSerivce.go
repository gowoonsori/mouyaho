package application

import (
	"likeIt/badge/infrastructure/badge"
	"likeIt/domain"
)

type LikeBadgeService struct {
	rr domain.ReactRepository
}

func (lbs LikeBadgeService) GetBadgeFile(userId domain.UserId, reqUrl string) []byte {
	//query string parsing
	urlInfo := CreateUrlInfoFromUrl(reqUrl)
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
	f, err := lbs.renderBadge(*urlInfo, isLike, likeCount)
	if err != nil {
		return []byte{}
	}

	return f
}

func (lbs LikeBadgeService) renderBadge(urlInfo UrlInfo, isLike bool, likeCount int) ([]byte, error) {
	bi := badge.NewLikeBadge(urlInfo.LikeIconColor, urlInfo.CountTextColor, urlInfo.ShareIconColor, urlInfo.BackgroundColor,
		likeCount, isLike, urlInfo.IsClear,
	)

	wr, err := badge.NewLikeBadgeWriter()
	if err != nil {
		return []byte{}, err
	}

	svg, err := wr.RenderBadge(bi)
	if err != nil {
		return []byte{}, err
	}

	return svg, nil
}
