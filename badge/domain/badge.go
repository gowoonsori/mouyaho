package domain

type BadgeUseCase interface {
	RenderBadgeHtml(badge Badge) []byte
}

type Badge struct {
	Id string
	githubIssue
	badgeInfo
}

func NewBadge(id, repo, title, bg, icon, react, text, edge string) *Badge {
	return &Badge{
		Id:          id,
		githubIssue: githubIssue{Repo: repo, Title: title},
		badgeInfo:   badgeInfo{BgColor: bg, IconColor: icon, ReactColor: react, TextColor: text, Edge: edge},
	}
}

type reactBadge struct {
	Badge
}

type voteBadge struct {
	Badge
}
