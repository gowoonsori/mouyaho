package domain

import (
	"likeIt/internal/badge"
	"regexp"
)

const (
	defaultBg     = "#ffffff"
	defaultBorder = "#111111"
	defaultIcon   = "#111111"
	defaultReact  = "#ff6767"
	defaultText   = "#111111"
	defaultShare  = "#111111"
	defaultEdge   = "flat"
)

//go:generate mockery --name BadgeService --case underscore --inpackage
type BadgeService interface {
	CreateHeartBadge(b Badge, userId uint64) *badge.HeartBadge
}

//go:generate mockery --name BadgeRepository --case underscore --inpackage
type BadgeRepository interface {
	GetReactionsBy(id BadgeId) []reaction
}

type Badge struct {
	BadgeId
	badgeInfo
	reactions []reaction
}

type BadgeId struct {
	Owner string
	Repo  string
	Title string
}

func CreateBadge(owner, repo, title, bg, border, icon, react, text, share, edge string, reactions []reaction) *Badge {
	r, _ := regexp.Compile("#[0-9a-zA-Z]")

	if bg == "" || !r.MatchString(bg) {
		bg = defaultBg
	}
	if border == "" || !r.MatchString(border) {
		border = defaultBorder
	}
	if icon == "" || !r.MatchString(icon) {
		icon = defaultIcon
	}
	if react == "" || !r.MatchString(react) {
		react = defaultReact
	}
	if text == "" || !r.MatchString(text) {
		text = defaultText
	}
	if share == "" || !r.MatchString(share) {
		share = defaultShare
	}
	if edge != "round" {
		edge = defaultEdge
	}

	return &Badge{
		BadgeId: BadgeId{Owner: owner, Repo: repo, Title: title},
		badgeInfo: badgeInfo{
			BgColor:     bg,
			BorderColor: border,
			IconColor:   icon,
			ReactColor:  react,
			TextColor:   text,
			ShareColor:  share,
			Edge:        edge,
		},
		reactions: reactions,
	}
}

func RenderBadgeHtml(hb badge.HeartBadge) (f []byte) {
	wr := badge.HeartBadgeWriter

	f, err := wr.ParseFile(hb)
	if err != nil {
		f = nil
	}
	return
}

func (b Badge) GetHeartReactions() []reaction {
	res := make([]reaction, len(b.reactions))
	for _, r := range b.reactions {
		if r.isHeart() {
			res = append(res, r)
		}
	}
	return res
}

func (b Badge) GetHeartReactionByUserId(id uint64) *reaction {
	for _, r := range b.reactions {
		if r.UserId == id && r.isHeart() {
			return &r
		}
	}
	return nil
}

func (b Badge) GetHeartCount() (cnt int) {
	for _, r := range b.reactions {
		if r.isHeart() {
			cnt++
		}
	}
	return
}
