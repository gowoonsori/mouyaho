package main

import (
	"likeIt/auth/github"
	"likeIt/config"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	http.HandleFunc("/auth", github.LoginHandler)
	http.HandleFunc("/auth/callback", github.CallbackHandler)
	
	_ = http.ListenAndServe(":"+config.Server.Port, nil)
}
