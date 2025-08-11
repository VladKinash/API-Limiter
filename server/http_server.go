package server

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received: %s %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func StartServer() {
	http.Handle("/", loggingMiddleware(http.HandlerFunc(rootHandler)))
	http.ListenAndServe(":8080", nil)
}
