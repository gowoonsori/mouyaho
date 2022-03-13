package interaface

import "likeIt/domain"

type BadgeHandler struct {
	lbs domain.BadgeService
}

func (bh *BadgeHandler) GetBadge(userId uint64, bd BadgeDto) (f []byte) {
	b := domain.CreateBadge(bd.Owner, bd.Repo, bd.Title, bd.BgColor, bd.BorderColor, bd.IconColor, bd.ReactColor, bd.TextColor, bd.ShareColor, bd.Edge, nil)
	hb := bh.lbs.CreateHeartBadge(*b, userId)
	f = domain.RenderBadgeHtml(*hb)

	return
}
