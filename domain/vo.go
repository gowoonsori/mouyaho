package domain

const (
	heartContent = "+1"
)

type badgeInfo struct {
	BgColor     string
	BorderColor string
	IconColor   string
	ReactColor  string
	Text        int
	TextColor   string
	ShareColor  string
	Edge        string //flat, round
	IsReact     bool
}
