package domain

type Color string

type badgeInfo struct {
	BgColor    Color
	IconColor  Color
	ReactColor Color
	TextColor  Color
	Edge       string
}

type githubIssue struct {
	Repo  string
	Title string
}
