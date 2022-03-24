package github

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"likeIt/config"
	"log"
	"net/http"
)

// get user token after login
func getUserToken(code, state string) string {
	//generate req body
	requestJSON, _ := json.Marshal(TokenRequest{
		ClientId:     config.Github.ClientId,
		ClientSecret: config.Github.ClientSecret,
		Code:         code,
		RedirectUrl:  config.Github.CallbackUrl,
		State:        state,
	})
	req, err := http.NewRequest("POST", tokenAPI, bytes.NewBuffer(requestJSON))
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
	var gts TokenResponse
	_ = json.Unmarshal(resBody, &gts)

	return gts.AccessToken
}
