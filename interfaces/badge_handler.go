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

type LikeBadgeHandler struct {
	badgeService domain.BadgeService
}

func NewLikeBadgeHandler(badgeService domain.BadgeService) *LikeBadgeHandler {
	return &LikeBadgeHandler{badgeService: badgeService}
}

func (l *LikeBadgeHandler) GetLikeBadge(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieId)
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:       cookieId,
			Value:      auth.NewSid(),
			Path:       "",
			Domain:     "",
			Expires:    time.Now().Add(time.Hour * 24 * 365),
			RawExpires: "",
			MaxAge:     0,
			Secure:     true,
			HttpOnly:   true,
			SameSite:   http.SameSiteDefaultMode,
			Raw:        "",
			Unparsed:   nil,
		}
		http.SetCookie(w, cookie)
	} else if cookie.Expires.Before(time.Now()) {
		cookie.Expires = time.Now().Add(time.Hour * 24 * 365)
	}

	b := l.badgeService.GetBadgeFile(domain.UserId(cookie.Value), r.RequestURI)
	if b == nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
	}

	w.Header().Add("Content-Type", "text/html")
	w.Write(b)
}
