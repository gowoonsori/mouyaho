package badge

import (
	"context"
	"fmt"
	"github.com/google/go-github/v43/github"
)

type IssueInfo struct {
	Id     string `json:"id"`
	Number int    `json:"number"`
	Title  string `json:"title"`
}

func GetIssues(repo, title string) []*github.Issue {
	c := github.NewClient(nil)
	issues, d, a := c.Search.Issues(context.Background(),
		fmt.Sprintf("\"%s\" type:issue in:title repo:%s", title, repo), &github.SearchOptions{Sort: "created", Order: "asc"})

	fmt.Println(d, a)
	return issues.Issues
}
