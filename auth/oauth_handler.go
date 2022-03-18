package auth

import (
	"fmt"
	"likeIt/env"
	"likeIt/session"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	authUrl := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", env.Config.Github.ClientId, env.Config.Github.CallbackUrl)
	http.Redirect(w, r, authUrl, http.StatusFound)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	//get access token
	code := r.URL.Query().Get("code")
	origin := r.URL.Query().Get("origin")
	token := getUserToken(code)

	//save token in session
	s, _ := session.Store.Get(r, session.Name)
	s.Values["token"] = token

	err := s.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, origin, http.StatusFound)
}
