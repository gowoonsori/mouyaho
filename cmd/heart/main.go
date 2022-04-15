package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"mouyaho/config"
	"mouyaho/internal/auth"
	"mouyaho/internal/badges"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()
	initMiddleware(r)
	initRoute(r)

	_ = http.ListenAndServe(config.Server.Port, r)
}

func initRoute(r *chi.Mux) {
	ghr := &badges.GithubRepository{}
	hs := badges.NewHeartsService(ghr)
	hh := badges.NewHeartsHandler(hs)

	r.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hi"))
		})
		r.Get("/badge", hh.HeartsBadgeHandler)
	})
	r.Route("/api", func(r chi.Router) {
		r.Get("/auth", auth.LoginHandler)
		r.Get("/auth/callback", auth.CallbackHandler)
		r.Get("/token", func(w http.ResponseWriter, r *http.Request) {
			c, _ := r.Cookie("mh_session")
			if c == nil {
				http.Error(w, "Bad Request: Empty cookie", http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			token := auth.DecryptCookie(*c)
			json.NewEncoder(w).Encode(struct {
				Token string `json:"token"`
			}{
				Token: token,
			})
		})
	})
}

func initMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(httprate.LimitByIP(10, 1*time.Minute))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           180,
	}))
}
