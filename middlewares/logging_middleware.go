package middlewares

import (
	goapi "github.com/carlosealves2/go-api"
	"log"
	"net/http"
)

func LoggingMiddleware(next goapi.HandlerFunc) goapi.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s", req.Method, req.URL.Path)
		next(w, req)
	}
}
