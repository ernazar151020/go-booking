package main

import (
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/ernazar151020/go-packages/config"
	"github.com/ernazar151020/go-packages/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// return mux

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteLog)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux

}
