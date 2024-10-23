package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func Authenticate(auth string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		basic := strings.Split(header, "Basic ")[1]
		if basic != base64.StdEncoding.EncodeToString([]byte(auth)) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}