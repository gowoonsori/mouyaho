package application

import (
	"errors"
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

func (ls likeService) Like(lr *dto.LikeRequest) (dto.LikeResponse, error) {
	l := ls.lr.FindByBadgeIdAndUserId(lr.BadgeId, lr.UserId)
	if l.IsEmpty() {
		return dto.LikeResponse{}, errors.New("존재하지 않는 badge id입니다.")
	}

	l, err := ls.lr.Save(l)
	if err != nil {
		return dto.LikeResponse{}, err
	}

	return dto.LikeResponse{
		Success: true,
		Like:    l.GetLikerCount(),
	}, nil
}

func (ls likeService) UnLike(lr *dto.LikeRequest) (dto.LikeResponse, error) {
	l := ls.lr.FindByBadgeIdAndUserId(lr.BadgeId, lr.UserId)
	if l.IsEmpty() {
		return dto.LikeResponse{}, errors.New("존재하지 않는 badge id입니다.")
	}

	err := ls.lr.DeleteById(l.Id)
	if err != nil {
		return dto.LikeResponse{}, err
	}

	return dto.LikeResponse{
		Success: true,
		Like:    l.GetLikerCount(),
	}, nil
}
