package github

import (
	"encoding/base64"
	"fmt"
	"likeIt/config"
	"likeIt/session"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.URL.Query().Get("src")
	if origin == "" {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// generate state
	state := base64.URLEncoding.EncodeToString([]byte(origin))
	authUrl := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&state=%s",
		config.Github.ClientId, config.Github.CallbackUrl, state)
	http.Redirect(w, r, authUrl, http.StatusFound)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {

	//get access token
	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	token := getUserToken(code, string(state))

	//save token in session
	c, _ := session.Store.Get(r, session.Name)
	c.Values["token"] = token

	err := c.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s, _ := base64.URLEncoding.DecodeString(state)
	origin := string(s)
	log.Println(origin)
	http.Redirect(w, r, origin, http.StatusFound)
}
