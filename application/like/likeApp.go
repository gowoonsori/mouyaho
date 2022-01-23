package like

import (
	"likeIt/domain/badge"
	"likeIt/domain/react"
	"likeIt/domain/user"
)

type serviceInterface interface {
	Like(url string, userId user.UserId) Info
	UnLike(url string, userId user.UserId) Info
}

type Service struct {
	br badge.Repository
	rr react.Repository
}

func (s *Service) Like(url string, userId user.UserId) Info {
	b,err := s.br.FindByUrl(url)
}

func (s *Service) UnLike(url string, userId user.UserId) Info {
	//TODO implement me
	panic("implement me")
}



