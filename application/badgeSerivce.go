package application

import (
	"likeIt/domain"
	"likeIt/infrastructure/badge"
)

type LikeBadgeService struct {
	rr domain.ReactRepository
}

func NewLikeBadgeService(rr domain.ReactRepository) *LikeBadgeService {
	return &LikeBadgeService{rr: rr}
}

func (lbs LikeBadgeService) GetBadgeFile(userId domain.UserId, reqUrl string) []byte {
	//query string parsing
	urlInfo := badge.CreateUrlInfoFromUrl(reqUrl)
	u := urlInfo.Url

	//like count get
	likeCount := lbs.rr.FindCountByBadgeId(domain.BadgeId(u))

	//isLike get
	il := lbs.rr.FindByBadgeIdAndUserId(domain.BadgeId(u), domain.UserId(userId))
	var isLike bool
	if il != nil {
		isLike = true
	}

	//file redering
	f, err := badge.GenerateLikeBadge(*urlInfo, isLike, likeCount)
	if err != nil {
		return []byte{}
	}

	return f
}

func (lbs LikeBadgeService) LikeBadge(userId domain.UserId, badgeId domain.BadgeId) *ReactDto {
	//userId/badgeId로 react 조회
	react := lbs.rr.FindByBadgeIdAndUserId(badgeId, userId)

	if react != nil && react.IsLike() {
		return nil
	}
	if react == nil {
		if _, err := lbs.rr.Save(domain.ByOn(userId, badgeId)); err != nil {
			return nil
		}
	}
	lbs.rr.UpdateLikeStatusById(react.Id(), true)
	//만일 like인데 react가 존재하고 like 삳태면 fail
	//like인데 react존재 안하거나 react추가
	//like인데 react가 존재하고 unlike상태면 상태 변경

	cnt := lbs.rr.FindCountByBadgeId(badgeId)
	//좋아요 수와 react상태 반환
	return &ReactDto{Count: cnt, IsLike: true}
}
func (lbs LikeBadgeService) ReactBadge(userId domain.UserId, badgeId domain.BadgeId, reactType string) {
	//userId/badgeId로 react 조회
	react := lbs.rr.FindByBadgeIdAndUserId(badgeId, userId)

	if react == nil {
		switch reactType {
		case "like":
			l, err := lbs.rr.Save(domain.ByOn(userId, badgeId))
		case "unlike":

		}
	} else {
		if react.IsLike() {

		}
	}
	//reactType으로 분리
	//만일 like인데 react가 존재하고 like 삳태면 fail
	//like인데 react존재 안하거나 react추가
	//like인데 react가 존재하고 unlike상태면 상태 변경

	//unlike인데 react가 존재하고 like이면 unlike 상태 변경
	//unlike인데 react가 존재하고 unlike이면 fail
	//unlike인데 react가 존재안하면 fail

	//좋아요 수와 react상태 반환
}
