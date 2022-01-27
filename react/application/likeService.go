package application

import (
	"likeIt/user/domain"
)

type serviceInterface interface {
	Like(url string, userId domain.UserId) Info
	UnLike(url string, userId domain.UserId) Info
}

//
//type Service struct {
//	br badge.Repository
//	rr react.Repository
//}
//
//func (s *Service) Like(url string, userId domain.UserId) Info {
//	b, err := s.br.FindByUrl(url)
//}
//
//func (s *Service) UnLike(url string, userId domain.UserId) Info {
//	//TODO implement me
//	panic("implement me")
//}
