package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		// call the controller
		next.ServeHTTP(w, r)

		diff := time.Since(now)

		log.Printf("[%s], path: %s, took: %s", r.Method, r.URL.Path, diff)
	})
}
