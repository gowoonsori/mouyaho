package application

import (
	"likeIt/application/dto"
	"likeIt/domain/badge"
	"likeIt/domain/like"
)

type badgeService struct {
	br badge.Repository
	lr like.Repository
}

type BadgeServiceInterface interface {
	GetBadge(lr *dto.LikeRequest) (dto.BadgeInfoResponse, error)
}

var _ BadgeServiceInterface = &badgeService{} //type 검사 (interface를 구현했는지 컴파일타임에 검사

func (bs badgeService) GetBadge(lr *dto.LikeRequest) (dto.BadgeInfoResponse, error) {
	b, err := bs.br.FindById(lr.BadgeId)
	if err != nil {
		return dto.BadgeInfoResponse{}, err
	}

	l := bs.lr.FindByBadgeIdAndUserId(b.Id, lr.UserId)
	if l.IsEmpty() {
		return dto.BadgeInfoResponse{
			Like:   b.LikeCount,
			IsLike: false,
		}, nil
	}

	return dto.BadgeInfoResponse{
		Like:   b.LikeCount,
		IsLike: true,
	}, nil
}
