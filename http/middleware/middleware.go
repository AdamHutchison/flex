package middleware

import "github.com/gorilla/mux"

func RegisterGlobalMiddleware(mux *mux.Router) {
	mux.Use(LoggingMiddleware)
}