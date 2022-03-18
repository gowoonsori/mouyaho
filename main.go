package main

import (
	"likeIt/auth"
	"likeIt/env"
	"net/http"
)

func main() {
	http.HandleFunc("/auth", auth.AuthHandler)
	http.HandleFunc("/auth/callback", auth.CallbackHandler)

	_ = http.ListenAndServe(":"+env.Config.Server.Port, nil)
}
