package github

import (
	"fmt"
	"likeIt/auth"
	"likeIt/config"
	"likeIt/session"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.URL.Query().Get("src")
	if origin == "" {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// generate state
	state := []byte(auth.EncryptAES([]byte(origin), []byte(config.App.StateKey)))
	authUrl := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&state=%s",
		config.Github.ClientId, config.Github.CallbackUrl, state)
	http.Redirect(w, r, authUrl, http.StatusFound)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	//get access token
	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	token := getUserToken(code, state)

	//save token in session
	c, _ := session.Store.Get(r, session.Name)
	c.Values["token"] = token

	err := c.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	origin := auth.DecryptAES([]byte(state), []byte(config.App.StateKey))
	http.Redirect(w, r, origin, http.StatusFound)
}
