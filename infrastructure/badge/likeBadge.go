package badge

import (
	"likeIt/infrastructure/internal/badge"
	"strconv"
)

func GenerateLikeBadge(urlInfo UrlInfo, isLike bool, likeCount int) (f []byte, err error) {
	bo := badge.NewBadgeOpts()
	if urlInfo.LikeIconColor != "" {
		bo.LeftIconColor(urlInfo.LikeIconColor)
	}
	if urlInfo.CountTextColor != "" {
		bo.TextColor(urlInfo.CountTextColor)
	}
	if urlInfo.ShareIconColor != "" {
		bo.RightIconColor(urlInfo.ShareIconColor)
	}
	if urlInfo.BackgroundColor != "" {
		bo.BackgroundColor(urlInfo.BackgroundColor)
	}
	bo.Text(strconv.Itoa(likeCount)).IsReact(isLike).IsClear(urlInfo.IsClear)
	b := badge.CreateBadgeFromOpts(bo)

	wr, err := badge.NewLikeBadgeWriter()
	if err != nil {
		return
	}

	f, err = wr.RenderBadgeFile(*b)
	if err != nil {
		return
	}

	return
}
