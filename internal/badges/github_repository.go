package badges

import (
	"context"
	"fmt"
	"github.com/google/go-github/v43/github"
)

var (
	githubAPI = "https://api.github.com"
)

type IssueInfo struct {
	Id     string `json:"id"`
	Number int    `json:"number"`
	Title  string `json:"title"`
}

type ReactionInfo struct {
	Id   int64 `json:"id"`
	User struct {
		Id   int64  `json:"id"`
		Name string `json:"login"`
	} `json:"user"`
	Content string `json:"content"`
}

type GithubRepository struct{}

func (gh GithubRepository) GetIssues(repo, title string) []*github.Issue {
	c := github.NewClient(nil)
	issues, d, a := c.Search.Issues(context.Background(),
		fmt.Sprintf("\"%s\" type:issue in:title repo:%s", title, repo), &github.SearchOptions{Sort: "created", Order: "asc"})

	fmt.Println(d, a)
	return issues.Issues
}