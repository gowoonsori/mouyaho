package domain

type badgeInfo struct {
	BgColor     string
	BorderColor string
	IconColor   string
	ReactColor  string
	TextColor   string
	ShareColor  string
	Edge        string //flat, round
	Type        string //좋아요, 투표
}

type githubIssue struct {
	Repo  string
	Title string
}
