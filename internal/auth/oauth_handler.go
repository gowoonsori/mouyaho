package auth

import (
	"encoding/base64"
	"fmt"
	"mouyaho/config"
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

	//save token in cookie
	c := CreateCookie(token)
	w.Header().Set("Set-Cookie", c.String())

	url, err := base64.StdEncoding.DecodeString(state)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, string(url), http.StatusFound)
}
