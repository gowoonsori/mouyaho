package oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"likeIt/env"
	"log"
	"net/http"
)

//func MainHandler(w http.ResponseWriter, r *http.Request) {
//	session, _ := session.store.Get(r, "session")
//	userInfo := session.Values["user"]
//	if userInfo == nil {
//		http.Redirect(w, r, "/login", http.StatusFound)
//	} else {
//		RenderHtmlTemplate(w, "main.html", userInfo)
//	}
//}

// redirect to github login form
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	authUrl := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", env.Config.Github.ClientId, env.Config.Github.CallbackUrl)
	http.Redirect(w, r, authUrl, http.StatusFound)
}

// get user token after login
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	// POST request to set URL
	requestJSON, _ := json.Marshal(GithubTokenRequest{
		ClientId:     env.Config.Github.ClientId,
		ClientSecret: env.Config.Github.ClientSecret,
		Code:         code,
		Scopes:       []string{"repo"},
	})
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)
	if err != nil {
		log.Panic("Error: Token Request Create Error")
	}

	// Get the Access token
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("Request failed")
	}

	// Response body converted to stringified JSON
	resBody, _ := ioutil.ReadAll(res.Body)
	var gts GithubTokenResponse
	json.Unmarshal(resBody, &gts)

	w.Write([]byte(gts.AccessToken))
}
