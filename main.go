package main

import (
	"likeIt/env"
	"likeIt/oauth"
	"net/http"
)

func main() {
	http.HandleFunc("/auth", oauth.AuthHandler)
	http.HandleFunc("/auth/callback", oauth.CallbackHandler)

	_ = http.ListenAndServe(":"+env.Config.Server.Port, nil)
}
