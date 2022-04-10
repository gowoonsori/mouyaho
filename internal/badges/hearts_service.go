package badges

import (
	"mouyaho/domain"
	"strconv"
	"strings"
)

type HeartsService struct {
	br domain.BadgeRepository
}

func (hs HeartsService) React(repo, token, issueNumber string) *domain.Reaction {
	r := strings.Split(repo, "/")
	if len(r) != 2 || r[0] == "" || r[1] == "" {
		return nil
	}
	n, err := strconv.Atoi(issueNumber)
	if err != nil {
		return nil
	}
	return hs.br.CreateHeartsInIssue(r[0], r[1], token, n)
}
