package application

import (
	"likeIt/application/dto"
	"likeIt/domain/like"
)

type likeService struct {
	lr like.Repository
}

type LikeServiceInterface interface {
	Like(lr *dto.LikeRequest) (dto.LikeResponse, error)
	UnLike(lr *dto.LikeRequest) (dto.LikeResponse, error)
}

var _ LikeServiceInterface = &likeService{} //type 검사 (interface를 구현했는지 컴파일타임에 검사

func (l likeService) Like(lr *dto.LikeRequest) (dto.LikeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l likeService) UnLike(lr *dto.LikeRequest) (dto.LikeResponse, error) {
	//TODO implement me
	panic("implement me")
}
