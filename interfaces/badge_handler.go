package interfaces

import (
	"likeIt/domain"
	"likeIt/infrastructure/auth"
	"net/http"
	"time"
)

var (
	cookieId = "glb_id"
)

type LikeBadge struct {
	badgeService domain.BadgeService
}

func (l *LikeBadge) GetLikeBadgeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieId)
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:       cookieId,
			Value:      auth.NewSid(),
			Path:       "",
			Domain:     "",
			Expires:    time.Time{},
			RawExpires: "",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   false,
			SameSite:   0,
			Raw:        "",
			Unparsed:   nil,
		}
		http.SetCookie(w, cookie)
	}

	b := l.badgeService.GetBadgeFile(domain.UserId(cookie.Value), r.RequestURI)
	if b == nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
	}

	w.Header().Add("Content-Type", "text/html")
	w.Write(b)
}
