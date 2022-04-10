package badges

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v43/github"
	"io/ioutil"
	"mouyaho/domain"
	"net/http"
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

func (gh GithubRepository) CreateHeartsInIssue(owner, repo, token string, issueNumber int) *domain.Reaction {
	u := fmt.Sprintf(githubAPI+"/repos/%v/%v/issues/%v/reactions", owner, repo, issueNumber)

	body, _ := json.Marshal(github.Reaction{Content: github.String("heart")})
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(body))
	if err != nil {
		return nil
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	m := &ReactionInfo{}
	resBody, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(resBody, m)
	return &domain.Reaction{
		Id:       m.Id,
		UserId:   m.User.Id,
		UserName: m.User.Name,
		Content:  m.Content,
	}
}
