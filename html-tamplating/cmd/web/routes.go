package main

import (
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/newmohib/go-sample-web-app/html-tamplating/pkg/config"
	"github.com/newmohib/go-sample-web-app/html-tamplating/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	// routs create with pat
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// create route with chi
	mux := chi.NewRouter()
	// middleware
	mux.Use(middleware.Recoverer)

	// router
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
