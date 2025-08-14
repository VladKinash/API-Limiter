package server

import (
	"fmt"
	"net/http"
	//"github.com/VladKinash/API-Limiter/middleware"
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
	//var ApiLimiter = middleware.NewLimiter(15, "Limiter")
	//var validApiKeys = []string{"APIKEY1", "APIKEY2"}
	http.Handle("/", loggingMiddleware(http.HandlerFunc(rootHandler)))
	http.ListenAndServe(":8080", nil)
}
