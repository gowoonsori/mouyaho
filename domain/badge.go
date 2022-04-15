package domain

import (
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

type BadgeService interface {
}

//go:generate mockery --name BadgeRepository --case underscore --inpackage
type BadgeRepository interface {
}

type Badge struct {
	BadgeId
	badgeInfo
	reactions []Reaction
}

type BadgeId struct {
	Repo  string
	Title string
}

func CreateBadge(repo, title, bg, border, icon, react, textClr, share, edge string, text int, isReact bool, reactions []Reaction) *Badge {
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
	if textClr == "" || !r.MatchString(textClr) {
		textClr = defaultText
	}
	if share == "" || !r.MatchString(share) {
		share = defaultShare
	}
	if edge != "round" {
		edge = defaultEdge
	}

	return &Badge{
		BadgeId: BadgeId{Repo: repo, Title: title},
		badgeInfo: badgeInfo{
			BgColor:     bg,
			BorderColor: border,
			IconColor:   icon,
			ReactColor:  react,
			Text:        text,
			TextColor:   textClr,
			ShareColor:  share,
			Edge:        edge,
			IsReact:     isReact,
		},
		reactions: reactions,
	}
}

func CreateBadgeFromDto(d BadgeDto) *Badge {
	return CreateBadge(d.Repo, d.Title, d.Bg, d.Border, d.Icon, d.React, d.TextColor, d.Share, d.Edge, d.Text, d.IsReact, nil)
}

func (b Badge) GetHeartReactions() []Reaction {
	res := make([]Reaction, len(b.reactions))
	for _, r := range b.reactions {
		if r.isHeart() {
			res = append(res, r)
		}
	}
	return res
}

func (b Badge) GetHeartReactionByUserId(id int64) *Reaction {
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
