package middleware

import (
	"log"
	"net/http"
)

func Hudai(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("I am hudai start")
		next.ServeHTTP(w, r)
		log.Println("I am hudai end")
	})
}
