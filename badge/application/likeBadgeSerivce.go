package application

import (
	"likeIt/badge/infrastructure/badge"
	"likeIt/domain"
	"net/url"
)

type LikeBadgeService struct {
	rr domain.ReactRepository
}

func (bu LikeBadgeService) GetBadge(userId string, reqUrl string) []byte {
	//query string parsing
	qs, err := parsingUrl(reqUrl)
	if err != nil {
		return nil
	}
	urlInfo := CreateUrlInfoFromMap(qs)
	u := urlInfo.Url

	//like count get
	likeCount := bu.rr.FindCountByBadgeId(domain.BadgeId(u))

	//isLike get
	il := bu.rr.FindByBadgeIdAndUserId(domain.BadgeId(u), domain.UserId(userId))
	var isLike bool
	if il != nil {
		isLike = true
	}

	//file redering
	f, err := bu.renderBadge(*urlInfo, isLike, likeCount)
	if err != nil {
		return []byte{}
	}

	return f
}

func (bu LikeBadgeService) renderBadge(urlInfo UrlInfo, isLike bool, likeCount int) ([]byte, error) {
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

func parsingUrl(reqUrl string) (map[string]string, error) {
	result := make(map[string]string)

	p, _ := url.Parse(reqUrl)
	rq, _ := url.QueryUnescape(p.RawQuery)
	m, err := url.ParseQuery(rq)
	if err != nil {
		return result, err
	}
	for k, v := range m {
		result[k] = v[0]
	}

	return result, nil
}
