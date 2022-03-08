package domain

type BadgeUseCase interface {
	RenderBadgeHtml(badge Badge) []byte
}

type Badge struct {
	Id string
	githubIssue
	badgeInfo
}

func NewBadge(id, repo, title, edge string, bg, icon, react, text Color) *Badge {
	return &Badge{
		Id:          id,
		githubIssue: githubIssue{Repo: repo, Title: title},
		badgeInfo:   badgeInfo{BgColor: bg, IconColor: icon, ReactColor: react, TextColor: text, Edge: edge},
	}
}
