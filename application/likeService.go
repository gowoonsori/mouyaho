package application

import (
	"likeIt/application/dto"
	"likeIt/domain/badge"
	"likeIt/domain/like"
)

type likeService struct {
	likeRepository like.LikeRepository
}

type LikeServiceInterface interface {
	Like(like *dto.LikeRequest) (dto.LikeResponse, error)
	Likes(likes []like.Like) (int, error)
	UnLike(like like.Like) error
	UnLikes(likes []like.Like) (int, error)
	GetLikeCountByBadge(badge *badge.Badge) (dto.LikeResponse, error)
	GetLikeByUser(userId uint64) (bool, error)
}

var _ LikeServiceInterface = &likeService{} //type 검사 (interface를 구현했는지 컴파일타임에 검사

func (l likeService) Like(lr *dto.LikeRequest) (dto.LikeResponse, error) {
	like := lr.CreateLike()

	err := l.likeRepository.Like(&like)
	if err != nil {
		return dto.LikeResponse{Success: false, Like: 0}, err
	}

	cnt, err := l.likeRepository.GetLikeCountByBadgeId(lr.BadgeId)
	if err != nil {
		return dto.LikeResponse{Success: false, Like: 0}, err
	}

	return dto.LikeResponse{Success: true, Like: cnt}, nil
}

func (l likeService) Likes(likes []like.Like) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (l likeService) UnLike(like like.Like) error {
	//TODO implement me
	panic("implement me")
}

func (l likeService) UnLikes(likes []like.Like) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (l likeService) GetLikeCountByBadge(badge *badge.Badge) (dto.LikeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l likeService) GetLikeByUser(userId uint64) (bool, error) {
	if like, err := l.likeRepository.GetLikeByUserId(userId); like == nil || err != nil {
		return false, err
	}
	return true, nil
}
