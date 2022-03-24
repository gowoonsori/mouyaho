package github

import (
	"fmt"
	"likeIt/auth"
	"likeIt/config"
	"likeIt/session"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.URL.Query().Get("url")
	if origin == "" {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// generate state
	state := []byte(auth.EncryptAES([]byte(origin), []byte(config.App.StateKey)))
	a := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&state=%s",
		authorizeAPI, config.Github.ClientId, config.Github.CallbackUrl, state)
	http.Redirect(w, r, a, http.StatusFound)
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

	url := auth.DecryptAES([]byte(state), []byte(config.App.StateKey))
	http.Redirect(w, r, url, http.StatusFound)
}
