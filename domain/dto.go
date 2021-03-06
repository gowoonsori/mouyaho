package domain

type BadgeDto struct {
	Repo        string `schema:"repo"`
	IssueTerm   string `schema:"issue-term"`
	Title       string `schema:"title"`
	Origin      string `schema:"origin"`
	Path        string `schema:"path"`
	URL         string `schema:"url"`
	Description string `schema:"description"`
	SpecTitle   string `schema:"spec-title"`
	Bg          string `schema:"bg"`
	Border      string `schema:"border"`
	Icon        string `schema:"icon"`
	React       string `schema:"react"`
	Text        int    `schema:"text"`
	TextColor   string `schema:"text-color"`
	Share       string `schema:"share"`
	IsReact     bool   `schema:"isReact"`
	Edge        string `schema:"edge"` //flat, round
}
