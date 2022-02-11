package badge

import (
	"likeIt/infrastructure/internal/badge"
	"strconv"
)

const (
	xRadius           = 15
	yRadius           = 15
	defaultLikeColor  = "red"
	defaultTextColor  = "black"
	defaultShareColor = "black"
	defaultBg         = "#eee"
	defaultText       = "0"
)

func GenerateLikeBadge(urlInfo UrlInfo, isLike bool, likeCount int) (f []byte, err error) {
	b := badge.NewBadge(urlInfo.LikeIconColor, strconv.Itoa(likeCount), urlInfo.CountTextColor, urlInfo.ShareIconColor, urlInfo.BackgroundColor,
		strconv.Itoa(xRadius), strconv.Itoa(yRadius), isLike, urlInfo.IsClear,
	)

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
