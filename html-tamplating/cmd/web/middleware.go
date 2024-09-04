package main

import (
	"fmt"
	"net/http"

	"github.com/choi2k/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})

}

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	//app.InProduction is get from main.go due to run the go run cmd/web/*.go

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler

}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	// this session is created in main.go as global variable for only main package
	return session.LoadAndSave(next)

}
