package session

import (
	"github.com/gorilla/sessions"
	"mouyaho/config"
)

var (
	Name  = "mouyaho_session"
	Store = sessions.NewCookieStore([]byte(config.App.SessionKey))
)

func init() {
	Store.Options.HttpOnly = true
	Store.Options.Secure = true
	Store.Options.MaxAge = 60 * 60 * 24 * 30 //1 month
}
