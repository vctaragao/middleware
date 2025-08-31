package middleware

import (
	"log"
	"net/http"
)

var middlewares = []Middleware{}

type Middleware func(next http.Handler) http.Handler

func Use(middleware ...Middleware) {
	middlewares = append(middlewares, middleware...)
}

func Chain(handler http.Handler) http.Handler {
	for _, m := range middlewares {
		handler = m(handler)
	}

	return handler
}

func Test(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("test middleware")

		next.ServeHTTP(w, r)

		log.Println("after test middleware")
	})
}

func Debug() {
	log.Println(middlewares)
}
