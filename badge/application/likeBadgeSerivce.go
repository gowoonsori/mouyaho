package application

import (
	"fmt"
	"likeIt/badge/infrastructure/badge"
	"likeIt/domain"
	"net/url"
	"strconv"
)

type LikeBadgeService struct {
	br domain.BadgeRepository
	rr domain.ReactRepository
}

func (bu LikeBadgeService) GetBadge(userId string, reqUrl string) *domain.Badge {
	//query string parsing
	qs, err := ParsingUrl(reqUrl)
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

	//badge file get
	id := domain.BadgeId(urlInfo.CreateBadgeUrl())
	b, err := bu.br.FindById(id)
	if err != nil {
		return nil
	}
	if b != nil {
		return domain.NewBadge(b.Id(), b.File())
	}

	//if file not exist, save after redering
	f, err := renderLikeBadge(*urlInfo, isLike, likeCount)
	if err != nil {
		return nil
	}
	b = domain.NewBadge(id, f)
	if _, err = bu.br.Save(b); err != nil {
		fmt.Println(err)
	}

	return b
}

func renderLikeBadge(urlInfo UrlInfo, isLike bool, likeCount int) ([]byte, error) {
	bi := &badge.LikeBadge{
		IsReact:         isLike,
		LikeIconColor:   urlInfo.LikeIconColor,
		CountText:       strconv.Itoa(likeCount),
		CountTextColor:  urlInfo.CountTextColor,
		ShareIconColor:  urlInfo.ShareIconColor,
		BackgroundColor: urlInfo.BackgroundColor,
		IsTransparency:  urlInfo.IsTransparency,
	}

	wr, err := badge.NewLikeBadgeWriter()
	if err != nil {
		panic(err)
	}

	svg, err := wr.RenderBadge(*bi)
	if err != nil {
		panic(err)
	}

	return svg, nil
}

func ParsingUrl(reqUrl string) (map[string]string, error) {
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
