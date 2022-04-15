package badges

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Issue_Search_Success(t *testing.T) {
	//given
	repo, title := "gowoonsori/blog-comments", "home/"
	gh := GithubRepository{}

	//when
	issues := gh.GetIssues(repo, title)

	//then
	assert.Equal(t, 3, len(issues))
}
