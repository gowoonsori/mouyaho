package badges

import (
	"mouyaho/domain"
)

type HeartsService struct {
	br domain.BadgeRepository
}

func NewHeartsService(repository domain.BadgeRepository) *HeartsService {
	return &HeartsService{br: repository}
}