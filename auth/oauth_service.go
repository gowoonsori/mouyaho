package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"likeIt/env"
	"log"
	"net/http"
)

// get user token after login
func getUserToken(code string) string {
	// generate state
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	//generate req body
	requestJSON, _ := json.Marshal(GithubTokenRequest{
		ClientId:     env.Config.Github.ClientId,
		ClientSecret: env.Config.Github.ClientSecret,
		Code:         code,
		RedirectUrl:  env.Config.Github.CallbackUrl,
		State:        state,
	})
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)
	if err != nil {
		log.Panic("Error: Token Request Create Error")
	}

	// Get the Access token
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("Request failed")
	}

	// Response body converted to stringified JSON
	resBody, _ := ioutil.ReadAll(res.Body)
	var gts GithubTokenResponse
	json.Unmarshal(resBody, &gts)

	return gts.AccessToken
}
