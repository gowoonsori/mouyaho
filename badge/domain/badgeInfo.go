package domain

import "strconv"

const (
	defaultLikeColor  = "red"
	defaultTextColor  = "black"
	defaultShareColor = "black"
	defaultBg         = "#eee"
	defaultText       = "0"
)

type Writer interface {
	RenderBadge(b BadgeInfo) ([]byte, error)
}

type BadgeInfo struct {
	isReact       bool
	likeIconColor string

	countText      string
	countTextColor string

	shareIconColor string

	backgroundColor string
	isClear         bool
}

func NewBadgeInfo(isReact bool, likeIconColor string, countText string, countTextColor string, shareIconColor string, backgroundColor string, isClear bool) *BadgeInfo {
	return &BadgeInfo{isReact: isReact, likeIconColor: likeIconColor, countText: countText, countTextColor: countTextColor, shareIconColor: shareIconColor, backgroundColor: backgroundColor, isClear: isClear}
}

func (b *BadgeInfo) Init() {
	if b.likeIconColor == "" {
		b.likeIconColor = defaultLikeColor
	}
	if b.shareIconColor == "" {
		b.shareIconColor = defaultShareColor
	}
	if b.countTextColor == "" {
		b.countTextColor = defaultTextColor
	}
	if b.backgroundColor == "" {
		b.backgroundColor = defaultBg
	}
	if _, err := strconv.Atoi(b.countText); err != nil {
		b.countText = defaultText
	}
}

func (b *BadgeInfo) IsReact() bool {
	return b.isReact
}

func (b *BadgeInfo) LikeIconColor() string {
	return b.likeIconColor
}

func (b *BadgeInfo) CountText() string {
	return b.countText
}

func (b *BadgeInfo) CountTextColor() string {
	return b.countTextColor
}

func (b *BadgeInfo) ShareIconColor() string {
	return b.shareIconColor
}

func (b *BadgeInfo) BackgroundColor() string {
	return b.backgroundColor
}

func (b *BadgeInfo) IsClear() bool {
	return b.isClear
}
