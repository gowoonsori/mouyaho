package main

import (
	"mouyaho/auth/github"
	"mouyaho/badge"
	"mouyaho/config"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	http.HandleFunc("/auth", github.LoginHandler)
	http.HandleFunc("/auth/callback", github.CallbackHandler)
	http.HandleFunc("/badge", badge.GetBadge)
	_ = http.ListenAndServe(":"+config.Server.Port, nil)
}
