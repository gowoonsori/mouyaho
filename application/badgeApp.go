package application

import (
	"likeIt/application/dto"
	"likeIt/domain/badge"
)

type badgeService struct {
	br badge.Repository
}

type BadgeServiceInterface interface {
	GetBadge(lr *dto.LikeRequest) (dto.LikeResponse, error)
}

var _ BadgeServiceInterface = &badgeService{} //type 검사 (interface를 구현했는지 컴파일타임에 검사

func (br badgeService) GetBadge(lr *dto.LikeRequest) (dto.LikeResponse, error) {
	//TODO implement me
	panic("implement me")
}
