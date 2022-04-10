package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"mouyaho/config"
	"mouyaho/internal/auth"
	"mouyaho/internal/badges"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(httprate.LimitByIP(10, 1*time.Minute))

	r.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hi"))
		})
		r.Get("/badge", badges.HeartsBadgeHandler)
	})
	r.Route("/api", func(r chi.Router) {
		r.Get("/auth", auth.LoginHandler)
		r.Get("/auth/callback", auth.CallbackHandler)
		r.Get("/token", func(w http.ResponseWriter, r *http.Request) {
			c, _ := r.Cookie("mh_session")
			w.Write([]byte(auth.DecryptCookie(*c)))
		})
	})

	_ = http.ListenAndServe(config.Server.Port, r)
}
