package github

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mouyaho/config"
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

	// Get the Access token
	res, err := http.Post(tokenAPI, "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		log.Panic("Request failed")
	}

	defer res.Body.Close()

	// Response body converted to stringified JSON
	resBody, _ := ioutil.ReadAll(res.Body)
	var gts TokenResponse
	_ = json.Unmarshal(resBody, &gts)

	return gts.AccessToken
}
