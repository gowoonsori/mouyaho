package application

import (
	"likeIt/application/dto"
	"likeIt/domain/badge"
)

type likeService struct {
	badgeRepository badge.Repository
}

type LikeServiceInterface interface {
	Like(lr *dto.LikeRequest) (dto.LikeResponse, error)
	UnLike(lr *dto.LikeRequest) (dto.LikeResponse, error)
	isLike(lr *dto.LikeRequest) bool
	GetLikeCountByBadge(lr *dto.LikeRequest) int
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

func (l likeService) GetLikeCountByBadge(lr *dto.LikeRequest) int {
	b, err := l.badgeRepository.FindById(lr.BadgeId)
	if err != nil {
		return 0
	}

	return b.GetLikeCount()
}

func (l likeService) isLike(lr *dto.LikeRequest) bool {
	b, err := l.badgeRepository.FindById(lr.BadgeId)
	if err != nil {
		return false
	}

	return b.IsLike(lr.UserId)
}
