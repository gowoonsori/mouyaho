package application

import (
	mocks "likeIt/badge/infrastructure"
	"likeIt/domain"
	"testing"
)

var (
	bs domain.BadgeService
)

func Test_RenderBadgeHtml(t *testing.T) {
	//given
	b := domain.CreateBadge("gowoonsori", "example", "7", "#fff", "#111", "#111", "#ff6767",
		"0", "#111", "round", nil)
	reactions := make(reaction)
	expectedBadge := domain.CreateBadge("gowoonsori", "example", "7", "#fff", "#111", "#111", "#ff6767",
		"0", "#111", "round", nil)
	br := &mocks.BadgeRepository{}
	br.On("GetReactionsByBadge", b).Return(b.Reactions)
	bs = NewLikeBadgeService(br)
}
