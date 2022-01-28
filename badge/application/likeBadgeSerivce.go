package application

import (
	"context"
	"likeIt/badge/infrastructure/badge"
	"likeIt/domain"
	"net/url"
	"strconv"
)

type LikeBadgeService struct {
	br domain.BadgeRepository
	rr domain.ReactRepository
}

func (bu LikeBadgeService) renderLikeBadge(reqUrl string, isLike bool, likeCount int) ([]byte, error) {
	qs, err := bu.ParsingUrl(reqUrl)
	if err != nil {
		return []byte{}, err
	}

	it, err := strconv.ParseBool(qs["transparency"])
	if err != nil {
		return []byte{}, err
	}
	
	bi := &badge.LikeBadge{
		IsReact:         isLike,
		LikeIconColor:   qs["like_color"],
		CountText:       strconv.Itoa(likeCount),
		CountTextColor:  qs["text_color"],
		ShareIconColor:  qs["share_color"],
		BackgroundColor: qs["bg"],
		IsTransparency:  it,
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

func (bu LikeBadgeService) GetBadge(ctx context.Context, reqUrl string) *domain.Badge {
	cu := ctx.Value("lid")
	userId := cu.(domain.UserId)

	//isLike 조회
	il := bu.rr.FindByBadgeIdAndUserId(1, userId)
	var isLike bool
	if il != nil {
		isLike = true
	}

	//like count 조회
	likeCount := bu.rr.FindCountByBadgeId(1)

	//badge render
	f, err := bu.renderLikeBadge(reqUrl, isLike, likeCount)
	if err != nil {
		panic(err)
	}

	return domain.NewBadge(reqUrl, userId, f, likeCount, isLike)
}

func (bu LikeBadgeService) ParsingUrl(reqUrl string) (map[string]string, error) {
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
