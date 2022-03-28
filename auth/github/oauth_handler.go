package github

import (
	"encoding/base64"
	"fmt"
	"mouyaho/config"
	"mouyaho/session"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.URL.Query().Get("url")
	if origin == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// generate state
	state := base64.StdEncoding.EncodeToString([]byte(origin))
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
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	url, err := base64.StdEncoding.DecodeString(state)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, string(url), http.StatusFound)
}
