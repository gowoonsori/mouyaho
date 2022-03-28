package middleware

import (
	"mouyaho/session"
	"net/http"
)

func MiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, _ := session.Store.Get(r, session.Name)
		token := s.Values["token"]

		//If token not exist, redirect to auth
		if token == nil {
			http.Redirect(w, r, "/auth", http.StatusFound)
		}

		r.Header.Set("Authorization", "Bearer "+token.(string))
		next(w, r)
	}
}
