package main

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// Log clients routes
		log.Printf("%s - %s\n", req.Method, req.RequestURI)

		next.ServeHTTP(res, req)
	})
}
