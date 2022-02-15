package main

import (
	"github.com/stretchr/testify/mock"
	"likeIt/application"
	"likeIt/domain/mocks"
	"likeIt/interfaces"
	"net/http"
)

func main() {
	rr := new(mocks.ReactRepository)
	rr.On("FindCountByBadgeId", mock.Anything).Return(0)
	rr.On("FindByBadgeIdAndUserId", mock.Anything, mock.Anything).Return(nil)

	bs := application.NewLikeBadgeService(rr)
	badge := interfaces.NewLikeBadgeHandler(bs)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/like-badge", badge.GetLikeBadge)

	http.ListenAndServe(":8080", mux)

}
