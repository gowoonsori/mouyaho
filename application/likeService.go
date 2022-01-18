package application

import (
	"likeIt/application/dto"
	"likeIt/domain/badge"
)

type likeService struct {
	badgeRepository badge.Repository
}

type LikeServiceInterface interface {
	Like(like *dto.LikeRequest) (dto.LikeResponse, error)
	Likes() (int, error)
	UnLike() error
	UnLikes() (int, error)
	isLike(lr *dto.LikeRequest) bool
	GetLikeCountByBadge() (dto.LikeResponse, error)
}

var _ LikeServiceInterface = &likeService{} //type 검사 (interface를 구현했는지 컴파일타임에 검사

func (l likeService) Like(lr *dto.LikeRequest) (dto.LikeResponse, error) {
	err := l.badgeRepository.AddLiker(lr.UserId)
	if err != nil {
		return dto.LikeResponse{Success: false, Like: 0}, err
	}

	b := l.badgeRepository.FindById(lr.BadgeId)
	if b == nil {
		return dto.LikeResponse{Success: false, Like: 0}, err
	}

	cnt := b.GetLikeCount()
	return dto.LikeResponse{Success: true, Like: cnt}, nil
}

func (l likeService) Likes() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (l likeService) UnLike() error {
	//TODO implement me
	panic("implement me")
}

func (l likeService) UnLikes() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (l likeService) GetLikeCountByBadge() (dto.LikeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l likeService) isLike(lr *dto.LikeRequest) bool {
	b := l.badgeRepository.FindById(lr.BadgeId)
	return b.IsLike(lr.UserId)
}
