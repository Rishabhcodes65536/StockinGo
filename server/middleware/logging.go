package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		start:=time.Now()
		next.ServeHTTP(w,r)
		log.Printf("Method: %s, URL: %s, Time: %s", r.Method, r.URL, time.Since(start))
	})
}