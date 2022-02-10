package interfaces

import (
	"encoding/json"
	"likeIt/domain"
	"net/http"
)

type LikeBadge struct {
	badgeService domain.BadgeService
}

func (l *LikeBadge) GetLikeBadgeHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")

	b := l.badgeService.GetBadgeFile(domain.UserId(user.(string)), r.RequestURI)
	if b == nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
	}

	w.Header().Add("Content-Type", "text/html")
	result, err := json.Marshal(b)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(http.StatusText(500)))
	}
	w.Write(result)
}
