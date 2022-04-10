package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mouyaho/config"
	"net/http"
	"time"
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
	req, err := http.NewRequest("POST", tokenAPI, bytes.NewBuffer(requestJSON))
	if err != nil {
		log.Panic("Request failed")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)
	var gts TokenResponse
	_ = json.Unmarshal(resBody, &gts)
	return gts.AccessToken
}
