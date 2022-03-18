package session

import (
	"github.com/gorilla/sessions"
	"likeIt/env"
)

var (
	Name  = "mouyaho_session"
	Store = sessions.NewCookieStore([]byte(env.Config.App.SessionKey))
)
