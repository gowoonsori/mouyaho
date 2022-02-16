package main

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/mock"
	"likeIt/application"
	"likeIt/domain"
	"likeIt/domain/mocks"
	"likeIt/interfaces"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error Loading .env file")
	}
}

func main() {
	env := os.Getenv("APP_ENV")
	port := os.Getenv("APP_PORT")

	// Repository 구현체 생성
	var rr domain.ReactRepository
	if env == "local" {
		mrr := new(mocks.ReactRepository)
		mrr.On("FindCountByBadgeId", mock.Anything).Return(0)
		mrr.On("FindByBadgeIdAndUserId", mock.Anything, mock.Anything).Return(nil)
		rr = mrr
	}

	// Service(UseCase) 구현체 생성
	bs := application.NewLikeBadgeService(rr)

	// Handler 생성
	badge := interfaces.NewLikeBadgeHandler(bs)

	// 라우터 생성
	mux := http.NewServeMux()
	mux.HandleFunc("/api/like-badge", badge.GetLikeBadge)

	http.ListenAndServe(port, mux)
}
