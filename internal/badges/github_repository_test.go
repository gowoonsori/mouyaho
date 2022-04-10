package badges

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Issue_Search_Success(t *testing.T) {
	//given
	repo, title := "gowoonsori/blog-comments", "home/"

	//when
	issues := GetIssues(repo, title)

	//then
	assert.Equal(t, 3, len(issues))
}

func Test_Issue_Reaction_Hearts_Success(t *testing.T) {
	//given
	owner, repo, token := "gowoonsori", "blog-comments", "ghu_MzsR78EuFl6j7ft1FJE3xoAo8vqzDl0Cm3qL"
	issueNumber := 1

	//when
	reaction := CreateHeartsInIssue(owner, repo, token, issueNumber)

	//then
	assert.NotNil(t, reaction)
	assert.Equal(t, "heart", reaction.Content)
	assert.Equal(t, "gowoonsori", reaction.User.Login)
}
